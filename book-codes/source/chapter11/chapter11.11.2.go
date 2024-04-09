package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// 全局变量
var count int

// 定义读写锁
var rLock sync.RWMutex

// 定义同步等待组
var wg sync.WaitGroup

// 数据读取函数
func read(i int) {
	// 加锁
	rLock.RLock()
	// 设置延时
	t := time.Duration(i*2) * time.Second
	time.Sleep(t)
	fmt.Printf("读操作，等待时间：%v 数据=%d\n", t.Seconds(), count)
	// 解锁
	rLock.RUnlock()
	wg.Done()
}

// 数据写入函数
func write(i int) {
	// 加锁
	rLock.Lock()
	// 写入数据
	count = rand.Intn(1000)
	// 设置延时
	t := time.Duration(i*2) * time.Second
	time.Sleep(t)
	fmt.Printf("写操作，等待时间：%v 数据=%d\n", t.Seconds(), count)
	// 解锁
	rLock.Unlock()
	wg.Done()
}

func main() {
	// 设置同步等待组
	wg.Add(6)
	// 设置随机数种子
	rand.Seed(time.Now().UnixNano())
	// 执行6次并发
	for i := 1; i < 4; i++ {
		go write(i)
	}
	for i := 1; i < 4; i++ {
		go read(i)
	}
	// 等待同步等待组执行并发
	wg.Wait()
}
