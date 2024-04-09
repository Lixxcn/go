package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Printf("当前时间：%v\n", time.Now())
	// 创建带缓存的通道
	c := make(chan int, 1)
	// 往通道写入数值
	c <- 10
	// 执行死循环
	for {
		// 关键字select从通道读取数据
		select {
		// 获取通道的数值
		case m := <-c:
			fmt.Printf("通道c的值：%v\n", m)
		// 通道为空的时候设置2秒超时，并终止循环
		case <-time.After(2 * time.Second):
			fmt.Printf("2秒后的时间：%v\n", time.Now())
			return
		}
	}
}
