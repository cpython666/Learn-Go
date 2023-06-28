package main

import "fmt"

func main() {
	var i, j int = 1, 2
	k := 3
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
}

//在函数中，简洁赋值语句 := 可在类型明确的地方代替 var 声明。
//函数外的每个语句都必须以关键字开始（var, func 等等），因此 := 结构不能在函数外使用。
