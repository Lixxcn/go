package main

import (
	"fmt"
	"sync"
	"time"
)

// 定义全局变量
// 定义同步等待组
var wg sync.WaitGroup

// 定义并发函数
func set_amap(m *sync.Map, b int) {
	// 参数m以指针接收者方式表示
	for i := 1; i < 5; i++ {
		m.Store("age", i+b)
		v, _ := m.Load("age")
		fmt.Printf("sync.Map的age数据：%v\n", v)
	}
	// 释放同步等待
	wg.Done()
}

func main() {
	// 记录程序开始时间
	start := time.Now()
	// 设置同步等待组
	wg.Add(2)
	var m sync.Map
	// 执行并发操作，sync.Map以指针形式作为参数传递
	go set_amap(&m, 0)
	go set_amap(&m, 10)
	// 等待同步等待组
	wg.Wait()
	// 记录程序结束时间并计算执行时间
	end := time.Now()
	consume := end.Sub(start).Seconds()
	fmt.Println("程序执行耗时(s)：", consume)
}
