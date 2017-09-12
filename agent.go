package yapool

import (
	"github.com/sirupsen/logrus"
	"golang.org/x/sync/syncmap"
	"net"
	"time"
)

type Agent struct {
	w          *WaitGroupWrapper
	CenterConn *syncmap.Map //key is ip(string)  ,  value is conn(*net.Conn)
}

func GetAgent(centerIPs []string) *Agent {
	centerConnMap := new(syncmap.Map)
	for _, centerIP := range centerIPs {
		centerConnMap.Store(centerIP, nil)
	}
	return &Agent{
		w:          new(WaitGroupWrapper),
		CenterConn: centerConnMap,
	}
}

func (a *Agent) Heartbeat(interval string) {
	a.CenterConn.Range(func(ip, conn interface{}) bool {
		go a.connectCenter(ip.(string), interval)
		return true
	})

}

func (a *Agent) SendMsgToCenter(l Level, msg ...interface{}) []error {
	errs := make([]error, 0)
	a.CenterConn.Range(func(ip, conn interface{}) bool {
		var tcpConn net.Conn
		if conn == nil {
			var err error
			tcpConn, err = net.Dial("tcp", ip.(string))
			if err != nil {
				errs = append(errs, err)
				return true
			}
		}
		byt, err := ConvertMsg(l, msg...).Encode()
		if err != nil {
			errs = append(errs, err)
			return true
		}
		byt = append(byt, '\n')
		//logrus.Info("send msg is ",string(byt))
		if _, err := tcpConn.Write(byt); err != nil {
			errs = append(errs, err)
			return true
		}
		return true
	})
	return errs
}

func (a *Agent) connectCenter(centerIP string, interval string) {
	var duration time.Duration
	var err error

	duration, err = time.ParseDuration(interval)
	if err != nil {
		logrus.Fatalf("cannot convert interval(%s) to time.Duration! [error] : %s", interval[0], err.Error())
	}
	//conn.SetWriteDeadline(time.Now().Add(duration))

	ticker := time.NewTicker(duration)
	for {
		select {
		case <-ticker.C:
			conn, err := net.Dial("tcp", centerIP)
			if err != nil {
				logrus.Errorf("TCP  dial  Center(%s)  error : %s", centerIP, err.Error())
				continue
			}
			a.loopHandle(conn)
		}
	}

}

func (a *Agent) loopHandle(conn net.Conn) {
	_, err := conn.Write([]byte(Heartbeat.String() + "\n"))
	if err != nil {
		if err.(net.Error).Timeout() {
			logrus.Warnf("heartbeat to center(%s) time out ! ", conn.RemoteAddr().String())
			a.w.Wrapper(func() {
				time.Sleep(time.Second * 1)
				a.loopHandle(conn)
			})
		} else {
			logrus.Errorf("send heartbeat to center(%s) error : %s", conn.RemoteAddr().String(), err.Error())

		}
	}
	//logrus.Info("send")

}
