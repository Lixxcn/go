package main

import (
	"fmt"
	"time"
)

func sent_data(ch, ch1 chan int) {
	for i := 0; i < 5; i++ {
		select {
		case ch <- i:
			fmt.Printf("ch写入数据：%v\n", i)
		case ch1 <- i:
			fmt.Printf("ch1写入数据：%v\n", i)
		}
	}
}

func get_data(ch, ch1 chan int) {
	for i := 0; i < 5; i++ {
		select {
		case i := <-ch:
			fmt.Printf("ch接收数据：%v\n", i)
		case i := <-ch1:
			fmt.Printf("ch1接收数据：%v\n", i)
		}
	}
}

func main() {
	// 创建通道
	ch := make(chan int)
	ch1 := make(chan int)
	go sent_data(ch, ch1)
	go get_data(ch, ch1)
	time.Sleep(5 * time.Second)
}
