package main

import (
	"fmt"
)

func main() {
	// 定义变量，不设初始值
	var a int
	a = 10
	// 定义变量并设置初始值
	var b int = 10
	// 批量定义变量，可根据情况是否设置初始值
	var (
		//c int
		//d string
		_ string
		e int = 10
	)
	// 多个变量同一数据类型
	var f, g int
	// 批量赋值
	f, g = 10, 10
	f = f + g
	// 定义变量并赋值，通过数值设置变量的数据类型
	h := 10
	fmt.Printf("定义变量，不设初始值：%d\n", a)
	fmt.Printf("定义变量并设置初始值：%d\n", b)
	fmt.Printf("批量定义变量，可根据情况是否设置初始值：%d\n", e)
	fmt.Printf("多个变量同一数据类型：%d\n", f)
	fmt.Printf("定义变量并赋值，通过数值设置变量的数据类型：%d\n", h)
}
