package factory

type Factory interface{
	GetCar()Car
}

type Car interface{
	Create()
}

//FactoryVendor -
type FactoryVendor struct{
	Factory
}

func NewFactoryVendor (f Factory) *FactoryVendor{
	return &FactoryVendor{
		Factory: f,
	}
}

func (v *FactoryVendor)Create(){
	v.GetCar().Create()
}

