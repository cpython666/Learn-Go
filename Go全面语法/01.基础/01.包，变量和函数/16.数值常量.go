package main

import "fmt"

const (
	// 将 1 左移 100 位来创建一个非常大的数字 ，即这个数的二进制是 1 后面跟着 100 个 0
	Big = 1 << 100
	// 再往右移 99 位，即 Small = 1 << 1，或者说 Small = 2
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}
func main() {
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
	//fmt.Println(needInt(Big))
}

//数值常量是高精度的 值。
//一个未指定类型的常量由上下文来决定其类型。
//再尝试一下输出 needInt(Big) 吧。
//（int 类型最大可以存储一个 64 位的整数，有时会更小。）
//（int 可以存放最大64位的整数，根据平台不同有时会更少。）
