package main

import (
	"fmt"
)

func myfun() {
	// 定义函数
	fmt.Printf("自定义函数")
}

func main() {
	// 定义函数变量
	var m func()
	// 将函数作为变量m的值
	m = myfun
	// 调用函数
	m()
}
