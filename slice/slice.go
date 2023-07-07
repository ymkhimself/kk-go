package main

import (
	"fmt"
)

type Stu struct {
	Name string
	Age  int
}

func f1() {
	var stus []Stu
	fmt.Println(stus == nil)           // true
	stus = append(stus, Stu{"asd", 1}) // asd,1
	fmt.Println(stus == nil)           // false
	fmt.Println(stus)
}

func f2() {
	stus := make([]Stu, 10)
	fmt.Println(stus)
}

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int64, n)
	fmt.Println(a)
}
