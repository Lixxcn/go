package main

import (
	"fmt"
)

func get_data(d interface{}) interface{} {
	fmt.Printf("参数值为：%v，数据类型：%T\n", d, d)
	return d
}

func main() {
	d := get_data("Tom")
	fmt.Printf("返回值为：%v，数据类型：%T\n", d, d)
	d1 := get_data(666)
	fmt.Printf("返回值为：%v，数据类型：%T\n", d1, d1)
}
