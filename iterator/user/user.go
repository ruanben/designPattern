package user

type User struct {
	Name string
	Age  int
}

type UserCollection struct{
	Users []*User
}


