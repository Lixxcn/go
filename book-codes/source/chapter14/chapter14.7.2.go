package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type PersonInfo struct {
	Name string `json:"name"`
	Age  int32  `json:"age"`
}

func main() {
	// 使用OpenFile()打开文件，设置O_CREATE模式，文件不存在则创建
	// 如果不想为OpenFile()设置参数，可以用Open()代替，实现效果一样
	f1, _ := os.OpenFile("output.json", os.O_RDWR|os.O_CREATE, 0755)
	// 定义结构体类型的切片
	var person []PersonInfo
	// 实例化结构体Decoder，实现数据读取
	data := json.NewDecoder(f1)
	// 将已读取的数据加载到切片person
	err := data.Decode(&person)
	// 如果err不为空值nil，则说明读取错误
	if err != nil {
		fmt.Printf("JSON读取失败：%v\n", err.Error())
	} else {
		fmt.Printf("JSON读取成功：%v\n", person)
	}
	// 关闭文件
	f1.Close()
}
