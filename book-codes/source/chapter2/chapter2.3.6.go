package main

import "fmt"

func main() {
	var a = 4
	var ptr *int
	fmt.Printf("a变量的内存地址：%v\n", &a)
	// 将变量a的内存地址赋给指针ptr
	ptr = &a
	fmt.Printf("指针ptr的内存地址为：%v\n", ptr)
	fmt.Printf("指针ptr的值为：%v\n", *ptr)
}
