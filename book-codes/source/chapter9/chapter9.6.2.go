package main

import (
	"fmt"
)

// 定义接口
type usb interface {
	connect()
}

// 定义结构体
type phone struct {
	name string
}

type camera struct {
	name string
}

// 定义结构体方法
func (p *phone) connect() {
	fmt.Printf("连接手机：%v\n", p.name)
}

func (c *camera) connect() {
	fmt.Printf("连接相机：%v\n", c.name)
}

func main() {
	var u usb
	p := phone{name: "华为"}
	c := camera{name: "索尼"}
	// 第一次使用类型断言
	u = &c
	if _, ok := u.(*camera); ok {
		fmt.Printf("执行相机连接\n")
		u.connect()
	}
	// 第二次使用类型断言
	u = &p
	if _, ok := u.(*phone); ok {
		fmt.Printf("执行手机连接\n")
		u.connect()
	}
}
