package main

import (
	"fmt"
	"time"
)

func main() {
	/* 时间戳 */
	var timestamp int64 = 1630315335
	// 时间戳转换结构体Time
	tm := time.Unix(timestamp, 0)
	fmt.Printf("时间戳转换结构体Time：%v\n", tm)
	// 时间戳转换字符串格式化
	tms := time.Unix(timestamp, 0).Format("2006-01-02 15:04:05")
	fmt.Printf("时间戳转换字符串格式化：%v\n", tms)

	/* 结构体Time */
	now := time.Now()
	// 结构体Time转换时间戳
	// 生成秒级时间戳
	fmt.Printf("结构体Time转换秒级时间戳：%v\n", now.Unix())
	// 生成纳秒级时间戳
	fmt.Printf("结构体Time转换纳秒级时间戳：%v\n", now.UnixNano())
	// 结构体Time转换字符串格式化
	tms1 := now.Format("2006-01-02 15:04:05")
	fmt.Printf("结构体Time转换字符串格式化：%v\n", tms1)

	/* 字符串格式化 */
	layout := "2006-01-02 15:04:05"
	timeStr := "2021-08-30 17:34:05.1099536"
	// 字符串格式化转换结构体Time
	// 函数Parse()是转换UTC时间格式
	timeObj, _ := time.Parse(layout, timeStr)
	fmt.Printf("字符串格式化转换结构体Time：%v\n", timeObj)
	// 字符串格式化转换结构体Time
	// 函数ParseInLocation()是转换当地时间格式
	timeObj1, _ := time.ParseInLocation(layout, timeStr, time.Local)
	fmt.Printf("字符串格式化转换结构体Time：%v\n", timeObj1)
	// 字符串格式化转换时间戳
	// 先转换结构体Time，再由结构体Time转换时间戳
	timeObj2, _ := time.ParseInLocation(layout, timeStr, time.Local)
	// 转换时间戳
	t1 := timeObj2.Unix()
	fmt.Printf("字符串格式化转换秒级时间戳：%v\n", t1)
	t2 := timeObj2.UnixNano()
	fmt.Printf("字符串格式化转换级纳秒时间戳：%v\n", t2)
}
