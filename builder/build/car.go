package build

//具体的产品，汽车
type Car struct{
	Wheel int     //汽车轮子数
	Color string   //汽车颜色
	Window int     //车窗数
}

//建造者接口
type CarBuild interface {
	AddWheel() int   //添加汽车轮子
	Paint() string //汽车喷绘
	AddWindow()int //添加汽车车窗
	CreateCar()Car //创建car
}

//指挥者：指挥具体的工厂工人
type Manufacturer struct{
	build CarBuild
}

func (m *Manufacturer) SetBuilder(b CarBuild){
	m.build = b
}

func (m *Manufacturer)BuildCar()Car{
	return m.build.CreateCar()
}
