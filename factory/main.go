package main

import (
	"designPattern/factory/factory"
	"designPattern/factory/factory/byd"
	"designPattern/factory/factory/tsl"
)

func main(){
	var Factory = factory.NewFactoryVendor(&byd.BydFactory{})

	Factory.Create()

	Factory = factory.NewFactoryVendor(&tsl.TslFactory{})
	Factory.Create()
}

