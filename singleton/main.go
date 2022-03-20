package singleton

import "sync"

type Instance struct{
}
var ins *Instance

func NewInstance2()*Instance{
	return ins
}

//NewInstance lazy mode
func NewInstance()*Instance{
	if ins == nil{
		ins = &Instance{}
	}
	return ins
}

var insMu sync.Mutex
func NewInstanceSafe()*Instance{
	insMu.Lock()
	defer insMu.Unlock()
	if ins == nil{
		ins = &Instance{}
	}
	return ins
}

func NewInstanceSafeCheck()*Instance{
	if ins == nil{
		insMu.Lock()
		defer insMu.Unlock()
		if ins == nil{
			ins = &Instance{}
		}
	}
	return ins
}

var once sync.Once
func NewInstanceSync()*Instance{
	once.Do(func(){
		ins = &Instance{}
	})
	return ins
}
