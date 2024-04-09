package main

import (
	"container/list"
	"fmt"
)

func main() {
	//// 定义列表变量
	//var l2 list.List
	//// 添加元素
	//l2.PushBack("Tom")
	//l2.PushBack("Tim")
	//l2.PushBack("Lily")
	//l2.PushBack("Mary")
	//// 遍历输出元素
	//for i := l2.Front(); i != nil; i = i.Next() {
	//	fmt.Printf("列表l2的元素是：%v\n", i.Value)
	//}

	// 定义列表变量
	var l2 list.List
	// 添加元素
	l2.PushBack("Tom")
	l2.PushBack("Tim")
	l2.PushBack("Lily")
	l2.PushBack("Mary")
	// 定义变量next
	var next *list.Element
	// 遍历输出元素
	for i := l2.Front(); i != nil; i = next {
		fmt.Printf("列表l2的元素是：%v\n", i.Value)
		// 设置变量next的值，用于执行下一次循环
		next = i.Next()
		// 删除元素
		l2.Remove(i)
	}
	fmt.Printf("列表l2的元素是：%v\n", l2)
}
