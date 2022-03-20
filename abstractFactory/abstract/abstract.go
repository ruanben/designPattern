package abstract

//Factory 工厂，生产汽车和自动驾驶系统
type Factory interface{
	GetCar()Car
	GetAutoPilot()AutoPiolot
}

//Car 汽车
type Car interface{
	Create()
}

//AutoPiolot -自动驾驶系统
type AutoPiolot interface{
	CreateAutoPiolotSystem()
}


