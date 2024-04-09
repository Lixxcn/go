package main

import (
	"fmt"
)

func main() {
	// 只定义
	var m1 = map[string]string{}
	m1["name"] = "Tom"
	fmt.Printf("集合m1：%v\n", m1)
	// 定义并赋值
	var m2 = map[string]string{"name": "Lily"}
	fmt.Printf("集合m2：%v\n", m2)
	// 使用make()函数定义
	m3 := make(map[string]string)
	m3["name"] = "Tim"
	fmt.Printf("集合m3：%v\n", m3)
}
