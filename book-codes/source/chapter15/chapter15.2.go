package main

import (
	"fmt"
	"time"
)

func main() {
	// 获取当前时间
	now := time.Now()
	fmt.Printf("当前时间：%v\n", now)
	// 获取当前时间的年
	year := now.Year()
	fmt.Printf("获取当前时间的年：%v\n", year)
	// 获取当前时间的月
	month := now.Month()
	fmt.Printf("获取当前时间的月，英文格式：%v\n", month)
	fmt.Printf("获取当前时间的月，数字格式：%v\n", int(month))
	// 获取当前时间的日
	day := now.Day()
	fmt.Printf("获取当前时间的日：%v\n", day)
	// 获取当前时间的小时
	hour := now.Hour()
	fmt.Printf("获取当前时间的小时：%v\n", hour)
	// 获取当前时间的分钟
	minute := now.Minute()
	fmt.Printf("获取当前时间的分钟：%v\n", minute)
	// 获取当前时间的秒
	second := now.Second()
	fmt.Printf("获取当前时间的秒：%v\n", second)
	// 获取当天是星期几
	wk := now.Weekday()
	fmt.Printf("获取当天是星期几：%v\n", wk)
}
