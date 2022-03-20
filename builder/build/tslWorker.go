package build

import "fmt"

//具体的构建者 工人
type TslWorker struct{

}

func (t *TslWorker)CreateCar()Car{
	wheelNum := t.AddWheel()
	windowNum := t.AddWindow()
	color := t.Paint()
	return Car{
		Wheel:  wheelNum,
		Color:  color,
		Window: windowNum,
	}
}

func (t *TslWorker) AddWheel()int{
	fmt.Println("add 4 wheel for tsl car")
	return 4
}

func (t *TslWorker) Paint()string{
	fmt.Println("paint red for tsl car")
	return "red"
}

func(t *TslWorker)AddWindow()int{
	fmt.Println("add 4 window for tsl car")
	return 4
}
