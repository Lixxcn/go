package main

import "fmt"

func main() {
	// 定义匿名结构体
	var p struct {
		name string
		age  int
	}
	// 使用匿名结构体并为成员赋值
	p.name = "Tom"
	p.age = 10
	fmt.Printf("匿名结构体的成员name：%v\n", p.name)
	fmt.Printf("匿名结构体的成员age：%v\n", p.age)

	// 定义匿名结构体并赋值
	p1 := struct {
		name string
		age  int
	}{
		name: "Tim",
		age:  20,
	}
	fmt.Printf("匿名结构体的成员name：%v\n", p1.name)
	fmt.Printf("匿名结构体的成员age：%v\n", p1.age)
}
