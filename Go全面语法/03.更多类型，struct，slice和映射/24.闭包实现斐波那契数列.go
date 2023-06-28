package main

import "fmt"

func fib() func() int {
	prev, curr := 0, 1
	return func() int {
		result := prev
		prev, curr = curr, prev+curr
		return result
	}
}
func main() {
	f := fib()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
