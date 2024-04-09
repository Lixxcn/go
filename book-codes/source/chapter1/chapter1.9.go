package main

import "fmt"

func main() {
	// 定义变量name, age, addr
	// 用于存储用户输入的数据
	var name, age, addr string
	// 输出操作提示
	fmt.Printf("请输入你的名字：\n")
	// 存储用户输入数据
	fmt.Scanln(&name)
	// 输出操作提示
	fmt.Printf("请输入你的年龄：\n")
	// 存储用户输入数据
	fmt.Scanln(&age)
	// 输出操作提示
	fmt.Printf("请输入你的居住地：\n")
	// 存储用户输入数据
	fmt.Scanln(&addr)
	// 输出用户输入的所有数据
	fmt.Printf("你的名字是：%v，年龄：%v，居住地：%v", name, age, addr)
}
