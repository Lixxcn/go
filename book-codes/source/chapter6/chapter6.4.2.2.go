package main

import (
	"container/list"
	"fmt"
)

func main() {
	// 定义列表变量
	var l2 list.List
	fmt.Printf("列表l2：%v\n", l2)
	// 添加元素
	element := l2.PushBack("abc")
	fmt.Printf("列表l2：%v\n", l2)
	// 删除元素
	l2.Remove(element)
	fmt.Printf("列表l2：%v\n", l2)
}
