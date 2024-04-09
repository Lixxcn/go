package main

import "fmt"

// 定义结构体
type mystruct struct {
	name string
}

// 将mystyle定义为mystruct类型
type mystyle mystruct

func main() {
	// 将s声明为mystyle类型
	var s mystyle
	// 设置属性name的值
	s.name = "Tom"
	// 输出变量s的数据
	fmt.Printf("变量s的数据：%v\n", s)
	// 输出变量s的数据类型
	fmt.Printf("变量s的数据类型：%T\n", s)
}
