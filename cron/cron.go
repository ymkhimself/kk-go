package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"time"
)

var a = 0

func main() {
	cron1 := cron.New()
	_, err := cron1.AddFunc("@every 5s", func() {
		a++
		fmt.Println(a)
	})
	if err != nil {
		fmt.Println(err.Error())
	}
	cron1.Start()
	time.Sleep(time.Second * 1000)
}
