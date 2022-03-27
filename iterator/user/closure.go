package user

//闭包方式来实现iterator
func(u *UserCollection)Iterator()func()(*User, bool){
	index := 0
	return func()(val *User, ok bool){
		if index < len(u.Users){
			val, ok = u.Users[index], true
			index ++
			return
		}
		return
	}
}
