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
	// 使用OpenFile()打开文件，设置O_TRUNC模式，每次写入将覆盖原有数据
	// 如果不想为OpenFile()设置参数，可以用Create()代替，实现效果一样
	f2, _ := os.OpenFile("output.json", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	// 创建PersonInfo类型的切片
	p := []PersonInfo{{"David", 30}, {"Lee", 27}}
	// 实例化结构体Encoder，实现数据写入
	encoder := json.NewEncoder(f2)
	// 将变量p的数据写入JSON文件
	// 数据写入必须使用文件内容覆盖，即设置os.O_TRUNC模式，否则导致内容错乱
	err := encoder.Encode(p)
	// 如果err不为空值nil，则说明写入错误
	if err != nil {
		fmt.Printf("JSON写入失败：%v\n", err.Error())
	} else {
		fmt.Printf("JSON写入成功\n")
	}
}
