package main

import "fmt"

type Person struct {
	// 定义结构体，作为查询结果
	Name    string
	Age     int
	Address string
}

func findData(person []*Person, name string, age int) {
	// 查询数据
	for _, data := range person {
		if data.Name == name && data.Age == age {
			fmt.Println(data)
			return
		}
	}
	fmt.Println("没有找到对应的数据")
}

func main() {
	list := []*Person{
		{Name: "Lily", Age: 23, Address: "CN"},
		{Name: "Tom", Age: 25},
		{Name: "Lily", Age: 30},
	}
	// 多条件查询
	findData(list, "Lily", 23)
}
