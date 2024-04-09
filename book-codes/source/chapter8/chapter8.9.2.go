package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	// 定义结构体
	// 成员infos为切片，切片元素为结构体
	Infos []struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	} `json:"infos"`
}

func main() {
	// 定义结构体类型的切片，并赋值
	s := []struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}{{Name: "Tom", Age: 15}, {Name: "Lily", Age: 20}}
	// 实例化结构体并赋值
	p := person{Infos: s}
	// 输出结构体p
	fmt.Printf("结构体p为：%v\n", p)
	// 将结构体p转换JSON字符串
	data, _ := json.Marshal(&p)
	// 输出JSON字符串
	fmt.Printf("JSON数据为：%v\n", string(data))
}
