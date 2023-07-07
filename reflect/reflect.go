package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string
}

func f1(a interface{}) {
	b, ok := a.(string)
	fmt.Println(b, ok)
}

func com() bool {
	a := make(map[float64]int)
	a[1.5]=10
	b := make(map[float64]int)
	b[1.5]=10
	return reflect.DeepEqual(a, b)
}

func main() {
	fmt.Println(com())
}
