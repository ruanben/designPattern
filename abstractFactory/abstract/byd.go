package abstract

import "fmt"

type BydFactory struct{

}

func NewBydFactory()*BydFactory{
	return &BydFactory{}
}

func(bf *BydFactory)GetCar()Car{
	return &BydCar{}
}

func (bf *BydFactory)GetAutoPilot()AutoPiolot{
	return &BydAutoPiolot{}
}


type BydCar struct{

}

func NewBydCar()Car{
	return &BydCar{}
}

func (b *BydCar)Create(){
	fmt.Println("Create byd car")
}

type BydAutoPiolot struct{

}

func NewBydAutoPiolot()*BydAutoPiolot{
	return &BydAutoPiolot{}
}

func (b *BydAutoPiolot)CreateAutoPiolotSystem(){
	fmt.Println("byd auto piolot system is apllo")
}



