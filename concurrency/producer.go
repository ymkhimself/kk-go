package main

import (
	"fmt"
	"time"
)

/**
生产者，消费者问题
 */

func main() {
	queue := make(chan int, 0)
	prucuder := func() {
		a := 0
		for {
			queue <- a
			fmt.Println("生产:",a)
			a++
			time.Sleep(500)
		}
	}
	consumer := func () {
		for{
			t := <-queue
			fmt.Println("消费:",t)
		}
	}
	go prucuder()
	go consumer()
	select {
	
	}
}
