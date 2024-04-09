package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// 定义集合
	var m1 = map[string]interface{}{}
	m1["name"] = "Tom"
	m1["age"] = 10
	fmt.Printf("m1的数据为：%v\n", m1)
	var m2 = map[string]interface{}{}
	m2["name"] = "Lily"
	m2["age"] = 20
	fmt.Printf("m2的数据为：%v\n", m2)
	// 定义切片
	var s1 = []map[string]interface{}{m1, m2}
	fmt.Printf("s1的数据为：%v\n", s1)
	// 定义集合，键为字符串类型，值为接口类型
	var m3 = map[string]interface{}{}
	m3["infos"] = s1
	data, _ := json.Marshal(&m3)
	fmt.Printf("JSON数据为：%v\n", string(data))
}
