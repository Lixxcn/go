package main

import (
	"fmt"
	"time"
)

func main() {
	// 获取当前时间
	now := time.Now()
	// 字符串格式化的模板为Go的出生时间2006年1月2号15点04分 Mon Jan
	// 24小时制
	t1 := now.Format("2006-01-02 15:04:05.000 Mon Jan")
	fmt.Printf("24小时制：%v\n", t1)
	// 12小时制
	t2 := now.Format("2006-01-02 03:04:05.000 PM Mon Jan")
	fmt.Printf("12小时制：%v\n", t2)
	// 时间显示格式为：年/月/日 时:分
	t3 := now.Format("2006/01/02 15:04")
	fmt.Printf("时间显示格式为：年/月/日 时:分：%v\n", t3)
	// 时间显示格式为：时:分 年/月/日
	t4 := now.Format("15:04 2006/01/02")
	fmt.Printf("时间显示格式为：时:分 年/月/日：%v\n", t4)
	// 时间显示格式为：年/月/日
	t5 := now.Format("2006/01/02")
	fmt.Printf("时间显示格式为：年/月/日：%v\n", t5)
	// 时间显示格式为：年-月-日
	t6 := now.Format("2006-01-02")
	fmt.Printf("时间显示格式为：年-月-日：%v\n", t6)
	// 时间显示格式为：日-月-年
	t7 := now.Format("02-01-2006")
	fmt.Printf("时间显示格式为：日-月-年：%v\n", t7)
	// 时间显示格式为：时:分:秒
	t8 := now.Format("15:04:05.000")
	fmt.Printf("时间显示格式为：时:分:秒：%v\n", t8)
	// 时间显示格式为：时-分-秒
	t9 := now.Format("15-04-05.000")
	fmt.Printf("时间显示格式为：时-分-秒：%v\n", t9)
}
