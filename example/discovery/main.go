package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"yapool"
)

func main() {

	logrus.Info("center discover the service")
	center := yapool.GetCenter(":9007")
	err := center.ReceiveWithFunc(func(msg *yapool.Msg) {

		fmt.Println("receive msg is ", msg.Level, msg.Msg)
	}, "6s")
	if err != nil {
		logrus.Error(err)
	}
	/*if err := center.Receive("6s");err != nil {
		logrus.Error(err)
	}*/
	select {}
}
