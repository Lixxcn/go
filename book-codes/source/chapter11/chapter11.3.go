package main

import (
	"fmt"
	"time"
)

func running() {
	// 循环5次
	for i := 0; i < 5; i++ {
		fmt.Println("tick", i)
		// 延时1秒
		time.Sleep(1 * time.Second)
	}
}

func main() {
	// 并发执行程序
	go running()
	// 主程序
	for i := 0; i < 5; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println("Waitting for you")
	}
}

//func running(name string) {
//	// 循环5次
//	for i := 0; i < 5; i++ {
//		fmt.Printf("tick %v, %v\n", i, name)
//		// 延时1秒
//		time.Sleep(1 * time.Second)
//	}
//}
//
//func main() {
//	// 并发执行程序
//	var name = "Tom"
//	go running(name)
//	// 主程序
//	for i := 0; i < 5; i++ {
//		time.Sleep(1 * time.Second)
//		fmt.Println("Waitting for you")
//	}
//}
