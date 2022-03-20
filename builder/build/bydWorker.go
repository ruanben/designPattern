package build

import "fmt"

//具体的构建者 工人
type BydWorker struct{

}

func (b *BydWorker)CreateCar()Car{
	wheelNum := b.AddWheel()
	color := b.Paint()
	windowNum := b.AddWindow()
	return Car{
		Wheel:  wheelNum,
		Color:  color,
		Window: windowNum,
	}
}

func (b *BydWorker) AddWheel()int{
	fmt.Println("add 4 wheel for byd car")
	return 4
}

func (b *BydWorker) Paint()string{
	fmt.Println("paint blue for byd car")
	return "blue"
}

func(b *BydWorker)AddWindow()int{
	fmt.Println("add 4 window for byd car")
	return 4
}
