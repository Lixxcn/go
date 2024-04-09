package main

import "fmt"

func main() {
	// 定义int类型的指针变量
	var pint *int
	fmt.Printf("指针值为：%v，空间地址：%v\n", pint, &pint)
	// 定义float64类型的指针变量
	var pfloat *float64
	fmt.Printf("指针值为：%v，空间地址：%v\n", pfloat, &pfloat)
	// 定义string类型的指针变量
	var pstr *string
	fmt.Printf("指针值为：%v，空间地址：%v\n", pstr, &pstr)
	// 定义bool类型的指针变量
	var pbool *bool
	fmt.Printf("指针值为：%v，空间地址：%v\n", pbool, &pbool)
	// 定义byte类型的指针变量
	var pbyte *byte
	fmt.Printf("指针值为：%v，空间地址：%v\n", pbyte, &pbyte)

	ptr := new(int)
	fmt.Printf("ptr指向的变量值为：%v，空间地址：%v\n", *ptr, &ptr)
}
