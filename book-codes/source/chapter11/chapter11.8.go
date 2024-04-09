package main

import "fmt"

func main() {
	// 创建容量大小为2的通道
	ch := make(chan int, 2)
	// 往通道写入数据
	ch <- 666
	// 关闭通道
	close(ch)
	// 输出通道的长度和容量
	fmt.Printf("通道长度：%v，容量：%v\n", len(ch), cap(ch))
	// 关闭通道后再次写入数据
	//ch <- 777
	// 关闭通道后获取数据
	fmt.Printf("通道数据：%v\n", <-ch)
	fmt.Printf("通道数据：%v\n", <-ch)
}
