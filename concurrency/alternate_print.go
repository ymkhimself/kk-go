package main

import (
	"fmt"
	"sync"
)

/**
两个协程交替打印
*/

/**
思路1：
两个协程，一个锁，一个共享变量

思路2：
两个协程，两个channel，通过channel来控制

思路3：
两个协程，一个channel
*/

func alternate1() {
	mutex := sync.Mutex{}
	count := 0
	group := sync.WaitGroup{}
	f1 := func() {
		for count < 100 {
			mutex.Lock()
			if count%2 == 1 {
				fmt.Println("奇数:", count)
				count++
			}
			mutex.Unlock()
		}
		group.Done()
	}
	f2 := func() {
		for count < 100 {
			mutex.Lock()
			if count%2 == 0 {
				fmt.Println("偶数:", count)
				count++
			}
			mutex.Unlock()
		}
		group.Done()
	}
	group.Add(2)
	go f1()
	go f2()
	group.Wait()
}

func alternate2() {
	c1 := make(chan int, 1)
	c2 := make(chan int, 1)

	group := sync.WaitGroup{}
	group.Add(2)
	go func() {
		for {
			t := <-c1
			if t > 100 {
				c2<-t  // 自己退出前，还是要通知一下对方，不然会死锁。
				break
			}
			fmt.Println("奇数:", t)
			t++
			c2 <- t
		}
		group.Done()
	}()
	go func() {
		for {
			t := <-c2
			if t > 100 {
				c1<-t
				break
			}
			fmt.Println("偶数:", t)
			t++
			c1 <- t
		}
		group.Done()
	}()
	c1<-1
	group.Wait()
	close(c1)
	close(c2)
}


// 使用一个channel，其实就是当锁用了
func alternate3() {

}

func main() {
	alternate2()
}
