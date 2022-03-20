package abstract

import "fmt"

//具体tsl工厂
type TslFactory struct{
}

func NewTslFactory()*TslFactory{
	return &TslFactory{}
}

func(bf *TslFactory)GetCar()Car{
	return &TslCar{}
}

func (bf *TslFactory)GetAutoPilot()AutoPiolot{
	return &TslAutoPiolot{}
}

//具体tsl 汽车产品
type TslCar struct{

}

func NewTslCar()Car{
	return &BydCar{}
}

func (t *TslCar)Create(){
	fmt.Println("Create tsl car")
}

//具体tsl 自动驾驶系统
type TslAutoPiolot struct{

}

func NewTslAutoPiolot()*TslAutoPiolot{
	return &TslAutoPiolot{}
}

func (t *TslAutoPiolot)CreateAutoPiolotSystem(){
	fmt.Println("tsl auto piolot system is FSD")
}
