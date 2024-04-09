package main

import "fmt"

func main() {
	var x, y = 8, 5
	fmt.Printf("加法运算符：%d\n", x+y)
	fmt.Printf("减法运算符：%d\n", x-y)
	fmt.Printf("乘法运算符：%d\n", x*y)
	fmt.Printf("除法运算符：%d\n", x/y)
	fmt.Printf("取模运算符：%d\n", x%y)
	x++
	fmt.Printf("幂运算符：%d\n", x)
	y--
	fmt.Printf("取整运算符：%d\n", y)
}
