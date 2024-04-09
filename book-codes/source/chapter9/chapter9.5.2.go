package main

import (
	"fmt"
)

func main() {
	// 定义切片变量s，切片元素为空接口
	// 设置切片的元素值
	s := []interface{}{1, "abc", 1.32}
	fmt.Printf("切片数据：%v\n", s)

	// 定义集合变量m
	m := map[string]interface{}{}
	// 设置集合的键值对
	m["name"] = "Tom"
	m["age"] = 10
	fmt.Printf("集合数据：%v\n", m)

	// 定义匿名结构体ss
	var ss struct {
		name interface{}
	}
	// 设置结构体成员的数值
	ss.name = "Mary"
	fmt.Printf("结构体ss的数据：%v，数据类型：%T\n", ss.name, ss.name)
	ss.name = 10
	fmt.Printf("结构体ss的数据：%v，数据类型：%T\n", ss.name, ss.name)
}
