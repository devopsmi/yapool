package yapool

import "testing"

func TestAgent_Heartbeat(t *testing.T) {
	localIP := []string{"localhost:9007"}
	agent := GetAgent(localIP)
	agent.Heartbeat("5s")
}

func TestAgent_SendMsgToCenter(t *testing.T) {
	localIP := []string{"localhost:9007"}
	agent := GetAgent(localIP)
	agent.SendMsgToCenter(Heartbeat, "send msg to center!")
}

func TestConnect_Center(t *testing.T) {
	localIP := []string{"localhost:9007"}
	agent := GetAgent(localIP)
	agent.connectCenter(localIP[0], "3s")

}
