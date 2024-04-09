package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(2 * time.Second)
	i := 0
	for {
		// 执行定时间隔
		t := <-ticker.C
		fmt.Printf("当前时间：%v\n", t)
		i++
		fmt.Printf("当前循环次数：%v\n", i)
		if i == 3 {
			// 重置定时间隔
			ticker.Reset(3 * time.Second)
		}
		// 输出5次就停止
		if i == 5 {
			// 停止计时器
			ticker.Stop()
			// 终止循环
			break
		}
	}
}
