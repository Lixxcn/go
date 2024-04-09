package main

import "fmt"

func myFunc() {
	// 定义延时执行的匿名函数
	defer func() {
		// 使用recover()捕捉异常
		if err := recover(); err != nil {
			// err不为空值，说明主动抛出异常
			fmt.Printf("捕捉异常：%v\n", err)
		} else {
			// err为空值，说明程序没有抛出异常
			fmt.Println("程序没有异常")
		}
	}()
	// 正常执行程序
	fmt.Println("程序正常运行")
	// 主动抛出异常
	panic("这是自定义异常")
}

func main() {
	// 调用函数
	myFunc()
}
