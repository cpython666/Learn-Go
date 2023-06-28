package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}

//switch 的 case 语句从上到下顺次执行，直到匹配成功时停止。
//（例如，
//switch i {
//case 0:
//case f():
//}
//在 i==0 时 f 不会被调用。）
