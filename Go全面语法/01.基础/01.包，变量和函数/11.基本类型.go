package main

import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}

//Go 的基本类型有
//bool
//string
//int  int8  int16  int32  int64
//uint uint8 uint16 uint32 uint64 uintptr
//byte // uint8 的别名
//rune // int32 的别名
//// 表示一个 Unicode 码点
//float32 float64
//complex64 complex128
//本例展示了几种类型的变量。 同导入语句一样，变量声明也可以“分组”成一个语法块。
//int, uint 和 uintptr 在 32 位系统上通常为 32 位宽，在 64 位系统上则为 64 位宽。 当你需要一个整数值时应使用 int 类型，除非你有特殊的理由使用固定大小或无符号的整数类型。
