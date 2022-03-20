package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

type inode interface {
	clone(string)inode
	getName()string
}

type file struct{
	name string
}

func (f file)clone(newName string) inode {
	clonef := f
	clonef.name = newName
	return clonef
}

func (f file)getName()string {
	return f.name
}

type folder struct{
	children []inode
	name string
}

func (f folder)clone(newName string) inode{
	cloneFolder := f
	cloneFolder.name = newName
	var cloneFolderCld []inode
	for _, i := range f.children{
		cloneFolderCld = append(cloneFolderCld, i.clone(i.getName() + "_copy"))
	}
	//cloneFolderCld = append(cloneFolderCld, f.children...)
	cloneFolder.children = cloneFolderCld
	return cloneFolder
}

func (f folder)getName()string {
	return f.name
}
type User struct{
	name string
}

func NewUser(name string)User{
	return User{name: name}
}

func (u *User)Copy(i inode, copyName string){
	copyFile := i.clone(copyName)
	fmt.Printf("user %v copy the %v, coped file %v", u.name, i, copyFile)
	fmt.Println()
}

func deepCopy(dst, src interface{}) error {
	var buf bytes.Buffer
	if err := gob.NewEncoder(&buf).Encode(src); err != nil{
		return err
	}
	return gob.NewDecoder(bytes.NewBuffer(buf.Bytes())).Decode(dst)
}

func main(){
	user := NewUser("helloWorld")
	fl1 := file{name: "test1.txt"}
	fl2 := file{name: "test2.txt"}
	fl3 := file{name: "test3.txt"}
	fd := folder{
		children: []inode{fl2, fl3},
		name:     "test",
	}
	user.Copy(&fl1, "test_copy.txt")
	user.Copy(&fd, "test_copy")
}






