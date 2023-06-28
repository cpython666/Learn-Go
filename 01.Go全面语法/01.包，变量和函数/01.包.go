package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("My favorite number is", rand.Intn(10))
}

//每个 Go 程序都是由包构成的。
//程序从 main 包开始运行。
//本程序通过导入路径 "fmt" 和 "math/rand" 来使用这两个包。
//按照约定，包名与导入路径的最后一个元素一致。例如，"math/rand" 包中的源码均以 package rand 语句开始。
