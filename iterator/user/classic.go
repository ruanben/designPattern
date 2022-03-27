package user

func (u *UserCollection)CreateIterator()UserIterator{
	return UserIterator{
		index: 0,
		users: u.Users,
	}
}
//type iterator interface {
//	hasNext() bool
//	getNext() *User
//}

type UserIterator struct{
	index int
	users []*User
}

func (u *UserIterator)HasNext()bool{
	if u.index < len(u.users){
		return true
	}
	return false
}

func (u *UserIterator)GetNext()*User{
	if u.HasNext() {
		user := u.users[u.index]
		u.index += 1
		return user
	}
	return nil
}
