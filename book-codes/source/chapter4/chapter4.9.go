package main

import (
	"fmt"
)

func main() {
	// 定义变量action，用于功能选择
	var action string
	// 定义变量data，存储当前输入的数据
	var d1, d2 float64
	// 定义变量opt，存储输入的运算符
	var opt string
	// 定义变量result，存储计算结果
	var result interface{}
	// 设置死循环，执行多次计算
	for {
		// 输出操作提示
		fmt.Printf("请输入选择，按1计算，按2退出：\n")
		// 存储用户输入数据
		fmt.Scanln(&action)
		// 进入计算操作
		if action == "1" {
			for i := 0; i < 2; i++ {
				// 输出操作提示
				fmt.Printf("请输入数字：\n")
				// 存储用户输入数据
				if i == 0 {
					fmt.Scanln(&d1)
				} else {
					fmt.Scanln(&d2)
				}
			}
			// 输出操作提示
			fmt.Printf("请输入运算法则，可选择+-*/：\n")
			// 存储用户输入数据
			fmt.Scanln(&opt)
			// 根据输入运算符执行对应计算
			if opt == "+" {
				result = d1 + d2
			}
			if opt == "-" {
				result = d1 - d2
			}
			if opt == "*" {
				result = d1 * d2
			}
			if opt == "/" {
				if d2 != 0.0 {
					result = d1 / d2
				} else {
					result = "除数为0无法计算"
				}
			}
			// 输出计算结果
			fmt.Printf("%v %v %v = %v\n", d1, opt, d2, result)
		}
		// 退出死循环，终止死循环
		if action == "2" {
			break
		}
	}
}
