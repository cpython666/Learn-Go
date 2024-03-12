// 空格分隔，行末不允许空格
package main

import "fmt"

// for循环输出
func print_num_1(n int) {
	for i := 1; i <= n; i++ {
		fmt.Print(i)
		if i != 10 {
			fmt.Print(" ")
		}
	}
	fmt.Println()
}

// 递归输出
func print_num_2(n int) {
	if n > 0 {
		print_num_2(n - 1)
		fmt.Print(n)
		if n != 10 {
			fmt.Print(" ")
		} else {
			fmt.Println()
		}
	}
}

// 切片和range方法
func print_num_3(n int) {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, num := range numbers {
		fmt.Print(num)
		if num != 10 {
			fmt.Print(" ")
		} else {
			fmt.Println()
		}
	}
}
func main() {
	print_num_1(10)
	print_num_2(10)
	print_num_3(10)
}
