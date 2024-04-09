package main

import (
	"fmt"
)

type queryKey struct {
	// 定义结构体，作为查询条件
	Name string
	Age  int
}

type Person struct {
	// 定义结构体，作为查询结果
	Name    string
	Age     int
	Address string
}

// 定义集合，key为结构体queryKey，value为结构体Person的指针
var mapper = make(map[queryKey]*Person)

// 定义函数，建立多条件查询
func buildIndex(person []*Person) {
	for _, p := range person {
		// 查询的组合键为Name、Age构建的结构体
		key := queryKey{
			Name: p.Name,
			Age:  p.Age,
		}
		mapper[key] = p
	}
	fmt.Printf("集合mapper是数据：%v\n", mapper)
}

// 定义函数，用于查询数据
func queryData(name string, age int) {
	// 实例化结构体queryKey
	key := queryKey{Name: name, Age: age}
	// 从集合mapper查询数据
	result, ok := mapper[key]
	// 输出查询结果
	if ok {
		fmt.Printf("查询结果：%v\n", result)
	} else {
		fmt.Println("没有找到对应的数据")
	}
}

func main() {
	// 定义切片变量list
	list := []*Person{
		{Name: "Lily", Age: 23, Address: "CN"},
		{Name: "Tom", Age: 25},
		{Name: "Lily", Age: 30},
	}
	// 多键索引查询数据
	buildIndex(list)
	queryData("Lily", 23)
}
