package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	b := a[3:]
	c := a[:4]
	d := make(chan int,10)
	d<-1
	fmt.Println(cap(b), cap(c), cap(d), len(d))
}
