package main

import "fmt"

// 定义结构体person
type person struct {
	name string
	age  int
}

func main() {
	// 实例化方法1
	// 实例化结构体person，生成实例化对象p
	p := person{name: "Tom", age: 18}
	// 由实例化对象p访问成员
	fmt.Printf("结构体成员name的值：%v\n", p.name)
	fmt.Printf("结构体成员age的值：%v\n", p.age)

	// 实例化方法2
	// 实例化结构体
	var p1 person
	// 对结构体成员进行赋值操作
	p1.name = "Tim"
	p1.age = 22
	// 由实例化对象p1访问成员
	fmt.Printf("结构体成员name的值：%v\n", p1.name)
	fmt.Printf("结构体成员age的值：%v\n", p1.age)

	// 实例化方法3
	// 使用new()实例化结构体
	p3 := new(person)
	// 对结构体成员进行赋值操作
	p3.name = "LiLy"
	p3.age = 28
	// 由实例化对象p3访问成员
	fmt.Printf("结构体成员name的值：%v\n", p3.name)
	fmt.Printf("结构体成员age的值：%v\n", p3.age)

	// 实例化方法4
	// 取结构体的地址实例化
	p4 := &person{}
	// 对结构体成员进行赋值操作
	p4.name = "Mary"
	p4.age = 16
	// 由实例化对象p4访问成员
	fmt.Printf("结构体成员name的值：%v\n", p4.name)
	fmt.Printf("结构体成员age的值：%v\n", p4.age)
}
