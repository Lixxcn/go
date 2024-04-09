package main

import (
	"fmt"
	"sync"
)

// 创建同步等待组对象
var wg sync.WaitGroup

// 定义函数，用于执行并发操作
func fun1() {
	for i := 1; i <= 3; i++ {
		fmt.Println("fun1。。i，", i)
	}
	// 代表完成并发，同步等待组的等待对象减一
	wg.Done()
}

// 定义函数，用于执行并发操作
func fun2() {
	for j := 1; j <= 3; j++ {
		fmt.Println("fun2。。j，", j)
	}
	// 代表完成并发，同步等待组的等待对象减一
	wg.Done()
}

func main() {
	// 设置同步等待组最大的等待数量
	wg.Add(2)
	// 执行并发
	go fun1()
	go fun2()
	fmt.Println("main进入阻塞状态。。等待并发程序结束。。")
	// 主程序进入阻塞状态，等待并发程序执行完成
	wg.Wait()
	fmt.Println("main解除阻塞。。")
}
