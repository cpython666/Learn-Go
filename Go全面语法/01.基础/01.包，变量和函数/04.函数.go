package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}

//函数可以没有参数或接受多个参数。
//在本例中，add 接受两个 int 类型的参数。
//注意类型在变量名 之后。（这种类型声明形式出现的原因可以去搜索一下。）
