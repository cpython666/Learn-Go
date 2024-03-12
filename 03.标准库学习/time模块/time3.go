package main

import (
	"fmt"
	"time"
)

func main() {
	layout := "2006/01/02 15:04:05"
	timeString := "2021/07/15 23:59:59"

	t, err := time.Parse(layout, timeString)
	if err != nil {
		fmt.Println("时间解析错误:", err)
		return
	}

	fmt.Println("解析后的时间:", t)

	layout = "2006-01-02 15:04:05"
	timeString = "2021-07-15 23:59:59"

	t, err = time.Parse(layout, timeString)
	if err != nil {
		fmt.Println("时间解析错误:", err)
		return
	}

	fmt.Println("解析后的时间:", t)
}
