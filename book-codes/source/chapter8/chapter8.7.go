package main

import (
	"fmt"
	"math/rand"
	"time"
)

type cat struct {
	// 定义结构体
	name   string
	weight int
	// 结构体成员为匿名结构体
	habit struct {
		ambient string
		style   string
	}
}

func get_cat() *cat {
	// 定义构造函数
	// 设置随机数的种子
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(10)
	// 定义变量，用于设置结构体的成员值
	var name, ambient, style string
	var weight int
	// 根据随机数设置变量值
	if n <= 5 {
		name = "tiger"
		weight = 500
		ambient = "山林"
		style = "独居"
	} else {
		name = "lion"
		weight = 300
		ambient = "草原"
		style = "群居"
	}
	// 实例化结构体
	c := cat{
		name:   name,
		weight: weight,
		// 匿名结构体实例化
		habit: struct {
			ambient string
			style   string
		}{ambient: ambient, style: style},
	}
	return &c
}

func main() {
	// 调用构造函数，获取结构体实例化对象
	c := get_cat()
	fmt.Printf("猫科动物为：%v\n", c.name)
	fmt.Printf("体重为：%v\n", c.weight)
	fmt.Printf("居住环境：%v\n", c.habit.ambient)
	fmt.Printf("生活方式：%v\n", c.habit.style)
}
