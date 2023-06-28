package main

import "fmt"

func main() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}

//此时你可以去掉分号，因为 C 的 while 在 Go 中叫做 for。
