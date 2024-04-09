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
	// 定义字符串，用于记录JSON数据
	var j string
	j = `{"infos":[{"name":"Tom","age":15},{"name":"Lily","age":20}]}`
	// 实例化结构体
	var p person
	// 将JSON字符串转为结构体p
	json.Unmarshal([]byte(j), &p)
	// 遍历输出结构体成员Infos的值
	// 遍历切片，切片元素为结构体
	for _, value := range p.Infos {
		fmt.Printf("获取Infos的值，名字为：%v\n", value.Name)
		fmt.Printf("获取Infos的值，年龄为：%v\n", value.Age)
	}

}
