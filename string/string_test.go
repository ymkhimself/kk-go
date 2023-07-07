package string

import (
	"fmt"
	"reflect"
	"testing"
)

func TestString(t *testing.T) {
	a := "asd余明款"
	fmt.Println(len(a))
	for _, v := range a {
		fmt.Println(reflect.TypeOf(v))
	}
}
