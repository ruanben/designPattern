package user
func (u *UserCollection)IteratorChan() <-chan *User{
	uc := make(chan *User)
	go func(){
		for _, val := range u.Users{
			uc <- val
		}
		close(uc)
	}()
	return uc
}
