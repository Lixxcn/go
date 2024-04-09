package main

import "fmt"

func main() {
	// 为变量myInt创建Int类型的内存地址
	myInt := new(int)
	// 给变量myInt赋值
	*myInt = 666
	fmt.Printf("myInt类型：%T，数值：%v，地址：%v\n", *myInt, *myInt, myInt)

	// 定义指针变量myPro
	var myPro *int
	// 输出指针变量的信息
	fmt.Printf("myPro类型：%T，数值：%v，地址：%v\n", myPro, myPro, myPro)
	// 定义变量num并赋值
	num := 777
	// 将变量num的内存地址赋予指针myPro
	myPro = &num
	// 输出指针变量的信息
	fmt.Printf("myPro类型：%T，数值：%v，地址：%v\n", *myPro, *myPro, myPro)
}
