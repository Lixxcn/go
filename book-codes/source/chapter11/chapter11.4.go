package main

import (
	"fmt"
	"time"
)

func main() {
	// 并发执行程序
	go func() {
		// 定义匿名函数
		for i := 0; i < 5; i++ {
			fmt.Println("tick", i)
			// 延时1秒
			time.Sleep(1 * time.Second)
		}
	}() // 小括号调用匿名函数
	// 主程序
	for i := 0; i < 5; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println("Waitting for you")
	}
}
