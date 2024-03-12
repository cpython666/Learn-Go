package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("返回当前的本地时间")
	fmt.Println(time.Now())
	fmt.Println("将时间格式化为指定的字符串格式")
	current := time.Now()
	fmt.Println(current.Format("2006-01-02 15:04:05"))
	fmt.Println("时间加减法：")
	current = time.Now()
	fmt.Println(current.Add(time.Hour)) // 在当前时间上加上一个小时
	fmt.Println(time.Since(current))    // 获取当前时间与指定时间之间的时间差
	fmt.Println(time.Now().Unix())
	fmt.Println(time.Now().UTC())
	layout := "2006-01-02 15:04:05"
	timeString := "2021-07-15 23:59:59"
	t, err := time.Parse(layout, timeString)
	if err != nil {
		fmt.Println("时间解析错误:", err)
		return
	}

	fmt.Println("解析后的时间:", t)

}
