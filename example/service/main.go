package main

import (
	"github.com/sirupsen/logrus"
	"time"
	"yapool"
)

func main() {

	logrus.Info("service start to running")

	agent := yapool.GetAgent([]string{"localhost:9007"})
	go agent.Heartbeat("2s")
	time.Sleep(time.Second * 3)
	errs := agent.SendMsgToCenter(yapool.Register, "go home")
	if len(errs) != 0 {
		for _, err := range errs {
			logrus.Error(err)
		}
	}
	logrus.Info("service is running")

	select {}
}
