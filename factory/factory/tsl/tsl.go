package tsl

import (
	"designPattern/factory/factory"
	"fmt"
)

type TslFactory struct{

}
func (t *TslFactory)GetCar()factory.Car{
	return &TslCar{}
}

type TslCar struct{
}
func (t *TslCar) Create(){
	fmt.Println("create tsl car")
}
