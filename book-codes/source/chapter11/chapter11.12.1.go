package main

import (
	"fmt"
	"sync"
	"time"
)

// 定义全局变量
// 定义互斥锁
var s sync.Mutex

// 定义同步等待组
var wg sync.WaitGroup

// 定义并发函数
func set_map(m map[string]int, b int) {
	for i := 1; i < 5; i++ {
		// 加锁处理
		s.Lock()
		m["age"] = i + b
		fmt.Printf("集合map的age数据：%v\n", m["age"])
		// 解锁处理
		s.Unlock()
	}
	// 释放同步等待
	wg.Done()
}

func main() {
	// 记录程序开始时间
	start := time.Now()
	// 设置同步等待组
	wg.Add(2)
	m := map[string]int{"age": 10}
	// 执行并发操作
	go set_map(m, 0)
	go set_map(m, 10)
	// 等待同步等待组
	wg.Wait()
	// 记录程序结束时间并计算执行时间
	end := time.Now()
	consume := end.Sub(start).Seconds()
	fmt.Println("程序执行耗时(s)：", consume)
}
