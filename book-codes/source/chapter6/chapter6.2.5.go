package main

import "fmt"

func main() {
	// 内置函数cap()获取切片容量
	// 内置函数len()获取切片长度
	s1 := make([]int, 3, 4)
	fmt.Printf("切片变量s1的值：%v\n", s1)
	fmt.Printf("切片变量s1的长度：%v\n", len(s1))
	fmt.Printf("切片变量s1的容量：%v\n", cap(s1))
	// 第一次添加元素
	s1 = append(s1, 10)
	fmt.Printf("切片变量s1的值：%v\n", s1)
	fmt.Printf("切片变量s1的长度：%v\n", len(s1))
	fmt.Printf("切片变量s1的容量：%v\n", cap(s1))
	// 第二次添加元素
	s1 = append(s1, 10)
	fmt.Printf("切片变量s1的值：%v\n", s1)
	fmt.Printf("切片变量s1的长度：%v\n", len(s1))
	fmt.Printf("切片变量s1的容量：%v\n", cap(s1))
}
