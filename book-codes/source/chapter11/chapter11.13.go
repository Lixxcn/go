package main

import (
	"fmt"
	"sync"
	"time"
)

// 创建同步等待组对象
var wg sync.WaitGroup

func producer(intChan chan int, exitChan chan bool) {
	// 往通道intChan写入数据，每两秒执行一次写入
	for i := 0; i < cap(intChan); i++ {
		fmt.Printf("厨师完成菜式%v的制作\n", i)
		intChan <- i
		time.Sleep(2 * time.Second)
	}
	// 往通道exitChan写入数据
	exitChan <- true
	fmt.Printf("厨师完成所有菜式\n")
	// 代表完成并发，同步等待组的等待对象减一
	wg.Done()
}

func consumer(intChan chan int, exitChan chan bool) {
	// 设置标签
fors:
	for {
		// 监听通道intChan和exitChan
		select {
		case v, _ := <-exitChan:
			if v {
				fmt.Printf("厨师下班，店铺不经营！\n")
				// 终止标签fors的循环
				// 不能直接使用break
				// 因为case语句里面使用break默认对case有效
				break fors
			}
		case v, ok := <-intChan:
			if ok {
				fmt.Printf("顾客吃了菜式%v！\n", v)
			}
		default:
			fmt.Printf("顾客等待中...\n")
			time.Sleep(3 * time.Second)
		}
	}
	// 代表完成并发，同步等待组的等待对象减一
	wg.Done()
}

func main() {
	// 设置同步等待组最大的等待数量
	wg.Add(2)
	// 定义通道变量intChan和exitChan
	intChan := make(chan int, 3)
	exitChan := make(chan bool)
	// 创建两个Goroutine，分别代表生产者和消费者
	go producer(intChan, exitChan)
	go consumer(intChan, exitChan)
	// 主程序进入阻塞状态，等待并发程序执行完成
	fmt.Printf("main进入阻塞状态..等待并发程序结束\n")
	wg.Wait()
	fmt.Printf("main解除阻塞\n")
}
