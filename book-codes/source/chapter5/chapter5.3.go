package main

import "fmt"

func main() {
	//var a int = 200
	//fmt.Printf("变量a的空间地址：%v\n", &a)
	//// 定义int类型的指针变量
	//var pint *int
	//fmt.Printf("指针值为：%v，空间地址：%v\n", pint, &pint)
	//pint = &a
	//fmt.Printf("指针值为：%v，空间地址：%v\n", pint, &pint)
	//fmt.Printf("指针值的值为：%v，空间地址：%v\n", *pint, &pint)

	var b int = 100
	var pint *int
	fmt.Printf("指针值的值为：%v，空间地址：%v\n", pint, &pint)
	// 将变量b的内存地址赋值给指针pint
	pint = &b
	fmt.Printf("指针值的值为：%v，空间地址：%v\n", *pint, &pint)
	// 通过取值操作符“*”修改变量b的值
	*pint = 666
	fmt.Printf("指针值的值为：%v，空间地址：%v\n", *pint, &pint)
}
