package main

import (
	"fmt"
	"reflect"
)

type Persons struct {
	N string `key:"Name" value:"李四"`
	A int    `key:"Age" value:"30"`
}

func main() {
	// 定义集合，存储结构体标签
	result := map[string]string{}
	// 实例化结构体
	p := Persons{N: "张三", A: 20}
	// 反射结构体
	vtp := reflect.TypeOf(p)
	// 遍历反射结构体变量的所有成员
	for i := 0; i < vtp.NumField(); i++ {
		// 获取每个成员的结构体成员类型
		vinfo := vtp.Field(i)
		// 输出结构体成员的tag
		tag_key := vinfo.Tag.Get("key")
		tag_value := vinfo.Tag.Get("value")
		result[tag_key] = tag_value
	}
	fmt.Printf("集合的数据：%v\n", result)
}
