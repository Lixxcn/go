package main

import (
	"fmt"
	"time"
)

func main() {
	// 获取当前时间
	now := time.Now()
	// 生成秒级时间戳
	t1 := now.Unix()
	// 生成纳秒级时间戳
	t2 := now.UnixNano()
	fmt.Printf("现在的秒级时间戳：%v，类型：%T\n", t1, t1)
	fmt.Printf("现在的纳秒级时间戳：%v，类型：%T\n", t2, t2)
}
