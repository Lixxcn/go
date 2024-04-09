package main

import "fmt"

func main() {
	// 定义一个空的字符串类型的切片指针
	var pslice []*string
	fmt.Printf("切片指针的元素：%v，内存地址：%v\n", pslice, &pslice)
	// 定义变量a、b、c并赋值
	var a, b string
	a, b = "a", "b"
	fmt.Printf("变量a、b的内存地址：%v、%v\n", &a, &b)
	// 使用内置函数append()将变量a、b、c的内存地址添加到切片指针
	pslice = append(pslice, &a)
	pslice = append(pslice, &b)
	fmt.Printf("切片指针的元素：%v\n", pslice)
	// 输出切片指针的元素所对应的数值
	// 使用取值操作符“*”从内存地址取值
	for _, k := range pslice {
		fmt.Printf("切片指针的元素所对应值：%v\n", *k)
	}
	// 从切片指针修改变量a的值，输出变量a
	*pslice[0] = "hello"
	fmt.Printf("修改后的变量值为：%v\n", a)
	// 修改变量b的值，输出切片指针的变量b的值
	b = "Golang"
	fmt.Printf("修改后的变量值为：%v\n", *pslice[1])
}
