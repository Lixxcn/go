package main

import (
	"fmt"
	"sort"
	"time"
)

// 定义结构体
type Memorandum struct {
	Date  int64
	Event string
}

// 自定义类型MemorandumSort
type MemorandumSort []Memorandum

func (m MemorandumSort) Len() int {
	return len(m)
}

func (m MemorandumSort) Less(i, j int) bool {
	// 按结构体成员Date排序，如果Date相同，按Event排序
	if m[i].Date < m[j].Date {
		return true
	} else if m[i].Date > m[j].Date {
		return false
	} else {
		return m[i].Event < m[j].Event
	}
}

func (m MemorandumSort) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func main() {
	// 创建当前时间的时间戳
	now := time.Now().Unix()
	// 实例化结构体Memorandum
	m1 := Memorandum{Date: now, Event: "学习Go语言"}
	m2 := Memorandum{Date: now + 7250, Event: "继续学习Go语言"}
	m3 := Memorandum{Date: now + 9070, Event: "晚了，洗洗睡吧，不然秃头了"}
	m4 := Memorandum{Date: now + 3460, Event: "休息了，顺便吃顿饭"}
	// 创建自定义类型MemorandumSort的实例化对象
	m := MemorandumSort([]Memorandum{m1, m2, m3, m4})
	// 排序
	sort.Stable(m)
	// 遍历输出自定义类型MemorandumSort的数据
	for _, v := range m {
		// 将时间戳转为字符串时间
		t := time.Unix(v.Date, 0).Format("2006-01-02 15:04:05")
		fmt.Printf("备忘时间：%v，你要做%v\n", t, v.Event)
	}
}
