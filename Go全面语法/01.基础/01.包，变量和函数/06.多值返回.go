package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}

//函数可以返回任意数量的返回值。
//swap 函数返回了两个字符串。
