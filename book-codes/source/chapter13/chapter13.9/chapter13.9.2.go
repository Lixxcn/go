package main

import (
	"fmt"
	"sort"
)

// 定义结构体
type Person struct {
	name string
	age  int
}

// 自定义数据类型
type PersonList []Person

// 排序规则：首先按年龄排序（由小到大）
// 年龄相同时按姓名进行排序（按字符串的自然顺序）
// 自定义sort.Interface的Len()
func (list PersonList) Len() int {
	return len(list)
}

// 自定义sort.Interface的Less()
func (list PersonList) Less(i, j int) bool {
	if list[i].age < list[j].age {
		return true
	} else if list[i].age > list[j].age {
		return false
	} else {
		return list[i].name < list[j].name
	}
}

// 自定义sort.Interface的Swap()
func (list PersonList) Swap(i, j int) {
	var temp Person = list[i]
	list[i] = list[j]
	list[j] = temp
}

func main() {
	// 实例化结构体
	p1 := Person{"Tom", 19}
	p2 := Person{"Hanks", 19}
	p3 := Person{"Amy", 19}
	p4 := Person{"Tom", 20}
	p5 := Person{"Jogn", 21}
	p6 := Person{"Mike", 23}
	// 结构体实例化对象写入切片
	pList := PersonList([]Person{p1, p2, p3, p4, p5, p6})
	// 对切片pList进行排序
	sort.Sort(pList)
	// 输出排序后的切片
	fmt.Println(pList)
	// Stable()比Sort()稳定
	sort.Stable(pList)
	// 输出排序后的切片
	fmt.Println(pList)
}
