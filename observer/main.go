package main

import (
	"github.com/ruanben/designPattern/observer/observ"
	"github.com/ruanben/designPattern/observer/subject"
)

func main(){
	obser1 := observ.NewWeChatUser("1")
	obser2 := observ.NewWeChatUser("2")

	subjt := subject.NewWeChatPublicAccount("水泥木屋")
	subjt.AddObserver(obser1)
	subjt.AddObserver(obser2)
	subjt.PublishArtical()
}


