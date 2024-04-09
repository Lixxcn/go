package main

import "fmt"

// 闭包 = 函数 + 引用环境
func adder() func(int) int {
	// 定义函数adder()，返回值为匿名函数func(int) int
	var x int = 10
	// 匿名函数作为函数返回值
	return func(y int) int {
		x += y
		return x
	}
}

func main() {
	// 函数adder()是一个闭包:
	// 函数adder()内部有变量x（引用环境）和匿名函数
	// 匿名函数引用了其外部作用域中的变量x
	// 在函数adder()的生命周期内 变量x一直有效
	f := adder()
	fmt.Println(f(10))
	fmt.Println(f(20))
	f1 := adder()
	fmt.Println(f1(2000))
	fmt.Println(f1(5000))
}
