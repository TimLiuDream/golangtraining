package main

import "fmt"

// non-empty interface

func Foo(x interface{}) {
	if x == nil {
		fmt.Println("empty interface")
		return
	}
	fmt.Println("non-empty interface")
}

func main() {
	var x *int = nil
	Foo(x)
}
