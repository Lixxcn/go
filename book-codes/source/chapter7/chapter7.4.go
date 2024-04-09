package main

import (
	"fmt"
)

func main() {

	//res := func(n1 int, n2 int) int {
	//	return n1 + n2
	//}(10, 30)
	//
	//fmt.Printf("函数执行结果为：%v\n", res)
	//

	// 将匿名函数赋给函数变量myfun
	myfun := func(n1 int, n2 int) int {
		return n1 - n2
	}

	// 变量myfun的数据类型是函数类型，可以由该变量完成函数调用
	res2 := myfun(10, 30)
	res3 := myfun(50, 30)
	fmt.Printf("匿名函数调用第一次：%v\n", res2)
	fmt.Printf("匿名函数调用第二次：%v\n", res3)
	fmt.Printf("函数变量myfun的数据类型：%T\n", myfun)
}
