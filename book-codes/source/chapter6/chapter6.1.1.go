package main

import "fmt"

func main() {
	// 定义长度为2的数组
	var s [2]string
	// 输出数组元素
	for i := 0; i < len(s); i++ {
		fmt.Printf("数组第%v个元素是：%v\n", i+1, s[i])
	}
	// 修改数组的元素值
	s[0] = "100"
	// 输出数组元素
	for i := 0; i < len(s); i++ {
		fmt.Printf("数组第%v个元素是：%v\n", i+1, s[i])
	}
}
