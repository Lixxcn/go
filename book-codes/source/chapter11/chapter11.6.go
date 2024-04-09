package main

import (
	"fmt"
	"time"
)

func Goroutine1(ch chan string) {
	fmt.Println("start goroutine1")
	// 数据写入通道，由Goroutine2()读取
	ch <- "goroutine2"
	fmt.Println("goroutine1 send channel: goroutine2")
	// 读取Goroutine2()写入的数据
	data := <-ch
	fmt.Printf("goroutine1 get channel: %v\n", data)
	// 数据写入通道，由主函数main()读取
	ch <- "Main goroutine"
}

func Goroutine2(ch chan string) {
	fmt.Println("start goroutine2")
	// 读取Goroutine1()写入的数据
	data := <-ch
	fmt.Printf("goroutine2 get channel: %v\n", data)
	// 数据写入通道，由Goroutine1()读取
	ch <- "goroutine1"
	fmt.Println("goroutine2 send channel: goroutine1")
}

func main() {
	// 构建通道
	ch := make(chan string)
	// 执行并发
	go Goroutine1(ch)
	// 执行并发
	go Goroutine2(ch)
	// 延时5秒，使Goroutine1()和Goroutine2()相互读写通道数据
	time.Sleep(5 * time.Second)
	// 读取Goroutine1()写入的数据
	data := <-ch
	fmt.Printf("main goroutine get channel: %v\n", data)
}
