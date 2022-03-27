package subject

import (
	"fmt"
	"github.com/ruanben/designPattern/observer/observ"
)

//被观察者 接口
type Subject interface {
	//增加一个观察者
	AddObserver(obs observ.Observer)
	//删除一个观察者
	DeleteObserver(obs observ.Observer)
	//通知所有观察者
	Notify()
}

//被观察者，微信公众号
type WeChatPublicAccount struct{
	AccountId string //公众号id
	observerList []observ.Observer
}

func NewWeChatPublicAccount(accountId string)*WeChatPublicAccount{
	return &WeChatPublicAccount{AccountId: accountId}
}

func(w *WeChatPublicAccount)AddObserver(o observ.Observer){
	w.observerList = append(w.observerList, o)
}

func(w *WeChatPublicAccount)DeleteObserver(o observ.Observer){
	observLength := len(w.observerList)
	for index, obs := range w.observerList{
		if obs.GetId() == o.GetId(){
			//将最后一个元素和当前要删除的元素换个位置
			w.observerList[observLength - 1], w.observerList[index] = w.observerList[index], w.observerList[observLength - 1]
			w.observerList = w.observerList[:observLength - 1]
			return
		}
	}
}

func (w *WeChatPublicAccount)PublishArtical(){
	fmt.Println("publish an artical")
	w.Notify()
}

func(w *WeChatPublicAccount)Notify(){
	for _, o := range w.observerList{
		o.Update(w.AccountId)
	}
}


