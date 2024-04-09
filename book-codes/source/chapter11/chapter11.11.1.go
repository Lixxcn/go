package main

import (
	"fmt"
	"sync"
	"time"
)

// 定义互斥锁Mutex的全局变量
var (
	myMutex sync.Mutex
)

func get_data(name string) {
	// 加锁处理
	myMutex.Lock()
	// 程序执行
	fmt.Printf("这是：%v\n", name)
	// 解锁处理
	myMutex.Unlock()
}

func main() {
	// 执行并发
	go get_data("get_data")
	// 加锁处理
	myMutex.Lock()
	// 程序执行
	fmt.Printf("这是：%v\n", "Main")
	for i := 0; i < 3; i++ {
		// 每一秒输出一行数据
		time.Sleep(1 * time.Second)
		fmt.Printf("等待时间：%v秒\n", i+1)
	}
	// 解锁处理
	myMutex.Unlock()
	// 等待延时，为了等待并发程序执行完成
	// 可以改为WaitGroup等待
	time.Sleep(2 * time.Second)
}
