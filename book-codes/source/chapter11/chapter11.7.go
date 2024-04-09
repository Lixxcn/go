package main

import "fmt"

func main() {
	// 创建一个3个元素缓冲大小的整型通道
	ch := make(chan int, 3)
	// 查看当前通道的大小
	fmt.Println(len(ch))
	// 发送3个整型元素到通道
	for i := 0; i < 3; i++ {
		ch <- i
	}
	// 查看当前通道的大小
	fmt.Println(len(ch))
	for i := 0; i < 3; i++ {
		fmt.Println(<-ch)
	}
	// 查看当前通道的大小
	fmt.Println(len(ch))
	// 查看当前通道的容量
	fmt.Println(cap(ch))
}
