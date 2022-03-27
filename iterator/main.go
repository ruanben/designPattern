package main

import (
	"fmt"
	"github.com/ruanben/designPattern/iterator/user"
)

func main(){
	user1 := &user.User{
		Name: "hello",
		Age:  20,
	}
	user2 := &user.User{
		Name: "world",
		Age:  30,
	}
	userColl := user.UserCollection{
		Users: []*user.User{user1, user2},
	}

	//iter := userColl.CreateIterator()
	//for iter.HasNext(){
	//	user := iter.GetNext()
	//	fmt.Println("user %+v", user)
	//}

	//iter := userColl.Iterator()
	//for {
	//	if val, ok := iter();ok{
	//		fmt.Println(val)
	//	}else{
	//		break
	//	}
	//}

	uc := userColl.IteratorChan()
	for v := range uc{
		fmt.Println(v)
	}
}