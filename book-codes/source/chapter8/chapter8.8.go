package main

import "fmt"

type person struct {
	// 定义结构体
	name   string
	weight int
}

func (p *person) get_name(name string) string {
	// 定义结构体方法
	return "My name is " + name
}

func (p *person) init(name string, weight int) {
	// 定义结构体方法，用于初始化结构体成员
	p.name = name
	p.weight = weight
}

func main() {
	// 实例化结构体
	p := person{}
	// 调用结构体方法，初始化成员值
	p.init("Tom", 99)
	fmt.Printf("结构体的成员name的值：%v\n", p.name)
	fmt.Printf("结构体的成员weight的值：%v\n", p.weight)
	// 调用结构体方法
	name := p.get_name(p.name)
	fmt.Printf("结构体方法的返回值：%v\n", name)
}
