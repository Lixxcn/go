package main

import (
	"fmt"
)

// 定义空接口
type empty interface{}

func main() {
	// 创建空接口变量
	var e empty
	fmt.Printf("空接口的数据：%v，数据类型：%T\n", e, e)
	// 定义字符串变量s
	s := "hello golang"
	// 将字符串变量赋给空接口变量
	e = s
	fmt.Printf("空接口的数据：%v，数据类型：%T\n", e, e)

	// 定义整型变量n
	n := 120
	// 将整型变量赋给空接口变量
	e = n
	fmt.Printf("空接口的数据：%v，数据类型：%T\n", e, e)

	// 定义布尔变量b
	b := true
	// 将布尔变量赋给空接口变量
	e = b
	fmt.Printf("空接口的数据：%v，数据类型：%T\n", e, e)
}
