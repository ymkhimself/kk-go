package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan int64, 1)
	go func(c1 chan int64) {
		for {
			select {
			case <-c1:
				fmt.Println("关了")
			default:
				fmt.Println("没关")
			}
		}
	}(c1)
	time.Sleep(time.Second * 2)
	close(c1)
}
