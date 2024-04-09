package main

import (
	"fmt"
	"time"
)

func main() {
	// 获取当前本地时间
	now := time.Now()
	fmt.Printf("当前本地时间：%v\n", now)
	/* Sub()计算两个时间差值 */
	layout := "2006-01-02 15:04:05"
	timeStr := "2021-07-30 17:34:05.1099536"
	// 字符串格式化转换结构体Time
	// 函数Parse()是转换UTC时间格式
	tp, _ := time.Parse(layout, timeStr)
	fmt.Printf("某个时间点：%v\n", tp)
	// 计算两个时间的差值
	r := now.Sub(tp)
	fmt.Printf("两个时间差值：%v\n", r)
	fmt.Printf("两个时间相差小时数：%v\n", r.Hours())
	fmt.Printf("两个时间相差分钟数：%v\n", r.Minutes())
	fmt.Printf("两个时间相差秒数：%v\n", r.Seconds())
	fmt.Printf("两个时间相差毫秒数：%v\n", r.Milliseconds())
	fmt.Printf("两个时间相差微秒数：%v\n", r.Microseconds())
	fmt.Printf("两个时间相差纳秒数：%v\n", r.Nanoseconds())
}
