package main

import "fmt"

func main() {
	var m1 = map[string]string{}
	m1["name"] = "Tom"
	m1["age"] = "20"
	m1["addr"] = "GZ"
	fmt.Printf("集合m1的数据：%v\n", m1)
	// 删除key=addr的数据
	delete(m1, "addr")
	fmt.Printf("集合m1的数据：%v\n", m1)
}
