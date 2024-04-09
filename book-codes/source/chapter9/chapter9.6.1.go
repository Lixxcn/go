package main

import (
	"fmt"
)

func get_data(d interface{}) interface{} {
	if a, ok := d.(string); ok {
		fmt.Printf("返回值ok为：%v，数据类型%T\n", ok)
		fmt.Printf("返回值a为：%v，数据类型：%T\n", a, a)
		fmt.Printf("参数值为：%v，数据类型：%T\n", d, d)
		return d
	} else if _, ok := d.(int); ok {
		fmt.Printf("参数值为：%v，数据类型：%T\n", d, d)
		return d
	}
	return "什么类型都不是"
}

func main() {
	d := get_data("Tom")
	fmt.Printf("返回值为：%v，数据类型：%T\n", d, d)
	d1 := get_data(666)
	fmt.Printf("返回值为：%v，数据类型：%T\n", d1, d1)
	d2 := get_data(true)
	fmt.Printf("返回值为：%v，数据类型：%T\n", d2, d2)
}
