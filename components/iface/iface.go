package main

import "fmt"

type MyStruct struct {
	Doer Doer
}

type Doer interface {
	Do() string
}

type Worker struct{}

func (w Worker) Do() string {
	return "I am working"
}

type Teacher struct{}

func (t Teacher) Do() string {
	return "I am teaching"
}

func main() {
	w := MyStruct{Doer: Teacher{}}
	fmt.Println(w.Doer.Do())
}
