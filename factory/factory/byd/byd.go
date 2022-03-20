package byd

import (
	"designPattern/factory/factory"
	"fmt"
)

type BydFactory struct{

}

func (b *BydFactory)GetCar()factory.Car{
	return &BydCar{}
}

type BydCar struct {
}

func(b *BydCar)Create(){
	fmt.Println("create byd car")
}
