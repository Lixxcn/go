package main

import "fmt"

type person struct {
	// 定义结构体
	string
	int
	float64
	bool
}

func main() {
	// 实例化结构体
	p := person{"Tim", 20, 171.1, true}
	// 访问匿名成员并输出
	fmt.Printf("结构体的匿名成员string的值：%v\n", p.string)
	fmt.Printf("结构体的匿名成员int的值：%v\n", p.int)
	fmt.Printf("结构体的匿名成员float64的值：%v\n", p.float64)
	fmt.Printf("结构体的匿名成员bool的值：%v\n", p.bool)
}
