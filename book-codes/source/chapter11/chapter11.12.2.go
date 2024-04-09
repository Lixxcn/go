package main

import (
	"fmt"
	"sync"
)

func main() {
	// 定义sync.Map类型的变量m
	var m sync.Map

	// Store()写入数据
	m.Store("name", "Tom")
	m.Store("age", 10)
	m.Store("address", "beijing")
	m.Store("vocation", "student")

	// Load()读取数据
	name, _ := m.Load("name")
	fmt.Printf("sync.Map的name数据：%v\n", name)
	age, _ := m.Load("age")
	fmt.Printf("sync.Map的age数据：%v\n", age)

	// Delete()删除数据
	m.Delete("address")
	fmt.Printf("sync.Map的数据：%v\n", m)

	// LoadAndDelete()读取并删除数据
	vocation, ok := m.LoadAndDelete("vocation")
	if ok {
		// 读取成功后输出数据
		fmt.Printf("sync.Map的vocation数据：%v\n", vocation)
		// 查看读取成功后是否删除数据
		fmt.Printf("sync.Map的数据：%v\n", m)
	}

	// LoadOrStore()读取或新增数据
	// 新增数据，如果key不存在，将参数key和value作为新的键值对写入
	live, ok := m.LoadOrStore("live", "BJ")
	fmt.Printf("sync.Map新增live数据：%v\n", live)
	fmt.Printf("sync.Map的数据：%v\n", m)
	// 读取数据，如果key已存在直接获取已有的value，参数value的值不起作用
	ages, ok := m.LoadOrStore("age", 15)
	fmt.Printf("sync.Map读取age数据：%v\n", ages)

	// 遍历输出数据
	m.Range(func(key, value interface{}) bool {
		fmt.Printf("sync.Map的key：%v\n", key)
		fmt.Printf("sync.Map的value：%v\n", value)
		return true
	})
}
