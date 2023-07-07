package main

import (
	"context"
	"fmt"
	"time"
)

func fun(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("fun done")
			return
		default:
			fmt.Println("fun")
			time.Sleep(time.Second)
		}

	}
}

func fun2(ctx context.Context) {
	c := time.After(time.Second * 3)
	for {
		select {
		case <-ctx.Done():
			fmt.Println("fun2 done")
			return
		case <-c:
			fmt.Println("fun2 到时间了")
			return
		default:
			fmt.Println("fun2")
			time.Sleep(time.Second)
		}
	}
}

var m map[int]func()

func main() {
	m = make(map[int]func())
	func() {
		ctx1 := context.Background()
		ctx1, cancel := context.WithCancel(ctx1)
		m[1] = cancel
		go fun(ctx1)
	}()
	func() {
		ctx2 := context.Background()
		ctx2, cancel2 := context.WithCancel(ctx2)
		go fun2(ctx2)
		m[2] = cancel2
	}()

	time.Sleep(5 * time.Second)
	m[1]()
	m[2]()
	fmt.Println(m)
	delete(m, 1)
	delete(m, 2)
	fmt.Println(m)
	for {

	}
}
