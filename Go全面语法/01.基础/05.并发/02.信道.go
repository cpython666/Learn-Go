package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // 将和送入 c
}
func main() {
	fmt.Println(5 / 2)
	s := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // 从 c 中接收
	fmt.Println(x, y, x+y)
}

//信道是带有类型的管道，你可以通过它用信道操作符 <- 来发送或者接收值。
//ch <- v    // 将 v 发送至信道 ch。
//v := <-ch  // 从 ch 接收值并赋予 v。
//（“箭头”就是数据流的方向。）
//和映射与切片一样，信道在使用前必须创建：
//ch := make(chan int)
//默认情况下，发送和接收操作在另一端准备好之前都会阻塞。这使得 Go 程可以在没有显式的锁或竞态变量的情况下进行同步。
//以下示例对切片中的数进行求和，将任务分配给两个 Go 程。一旦两个 Go 程完成了它们的计算，它就能算出最终的结果。
