package main

import (
	"fmt"
)

func main() {
	// 单个常量定义方式一
	const a int = 10
	// 单个常量定义方式二
	const b = 20
	// 多个常量定义方式一
	const (
		c int = 10
		d     = "golang"
	)
	// 多个常量定义方式二
	const e, f = true, 20
	fmt.Printf("单个常量定义方式一：%d\n", a)
	fmt.Printf("单个常量定义方式二：%d\n", b)
	fmt.Printf("多个常量定义方式一：%d\n", c)
	fmt.Printf("多个常量定义方式二：%d\n", f)
}
