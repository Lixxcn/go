package main

import "fmt"

func fibonacci(n int) int {
	// 定义递归函数
	if n < 2 {
		return n
	}
	// 调用自身，传入不同参数值
	return fibonacci(n-2) + fibonacci(n-1)
}

func main() {
	var i int
	// 调用函数fibonacci()
	for i = 0; i < 10; i++ {
		fmt.Printf("%d ", fibonacci(i))
	}
}
