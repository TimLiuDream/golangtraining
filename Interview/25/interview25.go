package main

import "fmt"

type User struct {
}
type MyUser1 User
type MyUser2 = User

func (i MyUser1) m1() {
	fmt.Println("MyUser1.m1")
}
func (i User) m2() {
	fmt.Println("User.m2")
}

// MyUser1.m1
// User.m2
func main() {
	var i1 MyUser1
	var i2 MyUser2
	i1.m1()
	i2.m2()
}
