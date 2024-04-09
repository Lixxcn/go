package main

import "fmt"

func main() {
	var s []int
	var ss = []int{1, 2}
	var sss []int = make([]int, 3)
	fmt.Printf("只定义：%v，内存地址为：%v\n", s, &s)
	fmt.Printf("定义并赋值：%v，内存地址为：%v\n", ss, &ss)
	fmt.Printf("使用make()函数定义：%v，内存地址为：%v\n", sss, &sss)

	//var ss = []int{1, 2}
	//fmt.Printf("切片变量ss的元素值为：%v\n", ss)
	//// 修改第一个元素值
	//ss[10] = 100
	//fmt.Printf("切片变量ss的元素值为：%v\n", ss)
}
