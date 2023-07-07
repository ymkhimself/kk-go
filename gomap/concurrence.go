package main

import (
	"fmt"
	"sync"
)

// 测试map的并发问题
var m map[int]int = make(map[int]int)

func conTest() {
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			m[i] = i
			fmt.Println(m[i])
			//fmt.Println(i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
