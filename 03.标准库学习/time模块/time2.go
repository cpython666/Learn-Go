package main

import (
	"fmt"
	"time"
)

func main() {
	time.Sleep(time.Duration(1)) // 使程序休眠指定的时间，duration 是一个 time.Duration 类型的变量，表示时间间隔
	fmt.Println("开始啦~")
	ticker := time.NewTicker(time.Second) // 创建一个定时器，每隔指定时间触发一次
	for range ticker.C {
		fmt.Println("111")
	}
}
