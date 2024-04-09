package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Printf("当前时间：%v\n", time.Now())
	// 延时1秒
	time.Sleep(time.Second)
	fmt.Printf("延时1秒后：%v\n", time.Now())
	// 延时5秒
	time.Sleep(5 * time.Second)
	fmt.Printf("延时5秒后：%v\n", time.Now())
	// 延时1分钟
	time.Sleep(time.Minute)
	fmt.Printf("延时1分钟后：%v\n", time.Now())
	// 延时2分钟
	time.Sleep(2 * time.Minute)
	fmt.Printf("延时2分钟后：%v\n", time.Now())
	// 延时1小时
	time.Sleep(time.Hour)
	fmt.Printf("延时1小时后：%v\n", time.Now())
}
