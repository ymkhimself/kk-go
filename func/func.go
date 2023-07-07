package main

import "fmt"

type A struct {
	Age int
}

func (a *A) f() {
	fmt.Println("f")
}

func (a A) f2() {
	fmt.Println("f2")
}

func main() {
	a := "asd"
	b := &a
	*b = "dsa"
	fmt.Println(a)
}
