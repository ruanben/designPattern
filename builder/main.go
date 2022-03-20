package main

import "designPattern/builder/build"

func main(){
	var manufacturer build.Manufacturer
	manufacturer.SetBuilder(&build.BydWorker{})
	manufacturer.BuildCar()

	manufacturer.SetBuilder(&build.TslWorker{})
	manufacturer.BuildCar()
}
