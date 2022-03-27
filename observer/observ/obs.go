package observ

import "fmt"

//观察者接口
type Observer interface {
	GetId()string
	Update(string)
}

//具体观察者，微信用户
type WeChatUser struct{
	UserId string
}

func NewWeChatUser(id string)*WeChatUser{
	return &WeChatUser{UserId: id}
}
func(u *WeChatUser) GetId()string{
	return u.UserId
}

func(u *WeChatUser)Update(observerId string){
	fmt.Println(fmt.Sprintf("user %v get a notify from observer: %v", u.UserId, observerId))
}