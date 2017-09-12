# yapool


## 安装
    go get github.com/CrocdileChan/yapool
    
    
## 使用    
正常情况下使用yapool<br>服务端监听本地的9007端口，并将超时时长设置为6秒。


    func main() {
    center := yapool.GetCenter(":9007")
    if err := center.Receive("6s");err != nil {
		log.Fatal(err)
    }
    select {}
    }


客户端往服务端发送心跳，心跳间隔时间为2秒

  
  
    func main() {
	   agent := yapool.GetAgent([]string{"localhost:9007"})
	   agent.Heartbeat("2s")
     select {}
    }
    
  
 需要往服务端发送除了心跳以外的其他讯息<br>服务端
 
 
     func main() {
	    center := yapool.GetCenter(":9007")
	    err := center.ReceiveWithFunc(func(msg *yapool.Msg) {
		  fmt.Println("receive msg is ", msg.Level, msg.Msg)
	    }, "6s")
	    if err != nil {
		    log.Fatal(err)
	    }
	    select {}
      }




客户端


    
    func main() {
          agent := yapool.GetAgent([]string{"localhost:9007"})
	        go agent.Heartbeat("2s")
	        errs := agent.SendMsgToCenter(yapool.Register, "go home")
	        if len(errs) != 0 {
		        for _, err := range errs {
			          log.Fatal(err)  
		        }
	        }

	        select {}
        }


