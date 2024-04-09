package main

import (
	"fmt"
	"strconv"
)

func myfun(name string, age int) (string, bool) {
	// 参数name和age
	// (string, bool)是返回值的数据类型
	var n string
	var b bool
	if name != "" {
		// 字符串拼接
		n = name + " is existence, age is " + strconv.Itoa(age)
		b = true
	} else {
		n = "name is not existence"
		b = false
	}
	// 返回值
	return n, b
}

func main() {
	// 调用函数，并设置返回值
	s, _ := myfun("Tom", 15)
	fmt.Println(s)
	// 调用函数，虽然有返回值，但函数外不需要使用
	myfun("Tom", 15)
}
