package main

import (
	"fmt"
	"time"
)

func main() {
	// 获取当前本地时间
	now := time.Now()
	fmt.Printf("当前本地时间：%v\n", now)
	/* Add()根据增量（时分秒）计算时间 */
	// 当前时间加1小时后的时间
	times1 := now.Add(time.Hour)
	fmt.Printf("1小时后的时间：%v\n", times1)
	// 当前时间加2小时后的时间
	times2 := now.Add(2 * time.Hour)
	fmt.Printf("2小时后的时间：%v\n", times2)
	// 当前时间加15分钟后的时间
	times3 := now.Add(15 * time.Minute)
	fmt.Printf("15分钟后的时间：%v\n", times3)
	// 当前时间加15分钟后的时间
	times4 := now.Add(-2 * time.Hour)
	fmt.Printf("2小时前的时间：%v\n", times4)

	/* AddDate()根据增量（年月日）计算时间 */
	// 当前时间加1年后的时间
	times6 := now.AddDate(1, 0, 0)
	fmt.Printf("1年后的时间：%v\n", times6)
	// 当前时间的2年前的时间
	times7 := now.AddDate(-2, 0, 0)
	fmt.Printf("2年前的时间：%v\n", times7)
	// 当前时间的3年2月10天后的时间
	times8 := now.AddDate(3, 2, 10)
	fmt.Printf("3年2月10天后的时间：%v\n", times8)
}
