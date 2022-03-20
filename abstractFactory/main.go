package main

import "designPattern/abstractFactory/abstract"

func main(){
	abstract.NewBydFactory().GetCar().Create()
	abstract.NewBydFactory().GetAutoPilot().CreateAutoPiolotSystem()

	abstract.NewTslFactory().GetCar().Create()
	abstract.NewTslFactory().GetAutoPilot().CreateAutoPiolotSystem()
}
