package main

import (
	"fmt"
)

type worker interface {
	work()
}

type person struct {
	name string
	worker
}

func printWorker() {
	var w worker = person{}
	fmt.Println(w)
}
