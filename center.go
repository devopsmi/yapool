package yapool

import (
	"bufio"
	"github.com/sirupsen/logrus"
	"net"
	"strings"
	"sync"
	"time"
)

const (
	Alive = "alive"
)

type Center struct {
	sync.RWMutex
	port          string
	db            DB
	agentTimerMap map[string]*time.Timer
}

func GetCenter(port string) *Center {
	return &Center{
		port: port,
		db:   GetDB(),
		//agentTimer: new(syncmap.Map),
		agentTimerMap: make(map[string]*time.Timer),
	}
}

func (c *Center) Receive(timeout string) error {
	l, err := net.Listen("tcp", c.port)
	if err != nil {
		return err
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			return err
		}
		go handleHeartbeat(conn)
		go c.handleAgent(conn, timeout)
	}
}

func (c *Center) ReceiveWithFunc(handleMsg func(msg *Msg), timeout string) error {
	l, err := net.Listen("tcp", c.port)
	if err != nil {
		return err
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			return err
		}

		go c.handleAgent(conn, timeout)

		msg, err := decodeMsgFromAgent(conn)
		if err != nil {
			logrus.Errorf("center(%s) get message from agent(%s) error : %s", conn.LocalAddr().String(), conn.RemoteAddr().String(), err.Error())
			continue
		}
		if msg != nil {
			go handleMsg(msg)
		}
	}

}

func (c *Center) handleAgent(conn net.Conn, timeout string) {

	agentID := strings.Split(conn.RemoteAddr().String(), ":")[0]
	duration, err := time.ParseDuration(timeout)
	if err != nil {
		logrus.Errorf("parse timeout(%s)  to duration error : %s", timeout[0], err.Error())
		return
	}

	c.RLock()
	timer, ok := c.agentTimerMap[agentID]
	c.RUnlock()
	if ok {
		if !timer.Stop() {
			<-timer.C
		}
		timer.Reset(duration)
		//logrus.Info("reset is ", timer)
	} else {
		c.Lock()
		c.agentTimerMap[agentID] = time.AfterFunc(duration, func() {
			c.expire(duration, agentID)
		})
		c.Unlock()
		//logrus.Info("timer is ", timer)
		c.db.Add(agentID, Alive)
		logrus.Infof("agent(%s) register to the center(%s)", agentID, conn.LocalAddr().String())
	}

}

func handleHeartbeat(conn net.Conn) {
	reader := bufio.NewReader(conn)
	line, err := reader.ReadString('\n')
	if err != nil {
		logrus.Errorf("center(%s) read HEARTBEAT bytes from agent(%s) error : %s", conn.LocalAddr().String(), conn.RemoteAddr().String(), err.Error())
		return
	}
	logrus.Info("receive heartbeat is ", line)
}

func (c *Center) expire(duration time.Duration, agent string) {
	c.db.Delete(agent)
	c.Lock()
	delete(c.agentTimerMap, agent)
	c.Unlock()
	logrus.Warnf("agent(%s) die!", agent)
}

func decodeMsgFromAgent(conn net.Conn) (*Msg, error) {
	reader := bufio.NewReader(conn)
	lineByt, err := reader.ReadBytes('\n')
	if err != nil {
		return nil, err
	}
	if IsLevel(strings.Trim(string(lineByt), "\n")) {
		return nil, nil
	}
	msg, err := Decode(lineByt)
	if err != nil {
		return nil, err
	}
	return msg, nil
}
