package main

import (
	"fmt"
	"testing"
)

func TestSlice2(t *testing.T) {
	e := []int32{1, 2, 3}
	fmt.Println("cap of e before:", cap(e))
	e = append(e, 4, 5, 6, 7)
	fmt.Println("cap of e after:", cap(e))
}

func TestSliceAppend(t *testing.T) {
	// 操作的全是底层那一个数组，没有扩容
	arr1 := []int{1, 2, 3, 4, 5, 6}
	arr2 := arr1
	arr2 = append(arr2[:2], arr2[3:]...)
	fmt.Println(arr1)
	fmt.Println(arr2)
}
