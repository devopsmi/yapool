package yapool

import "testing"

func TestCenter_Receive(t *testing.T) {
	center := GetCenter(":9007")
	center.Receive("5s")
}

func TestCenter_ReceiveWithFunc(t *testing.T) {
	center := GetCenter(":9007")
	center.ReceiveWithFunc(func(msg *Msg) {
		t.Logf("receive msg level (%s),msg content is (%s) ", msg.level.String(), msg.msg[0].(string))
	}, "5s")
}

func TestHandle_Agent(t *testing.T) {
	center := GetCenter(":9007")
	center.handleAgent("agent", "5s")
	center.agentTimer.Range(func(key, value interface{}) bool {
		t.Logf("key is %s ,value is %s", key, value)
		return true
	})
}
