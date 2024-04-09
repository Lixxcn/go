package main

import (
	"fmt"
)

//	func myfun(numbers ...int) {
//		for _, k := range numbers {
//			fmt.Printf("参数值为：%v\n", k)
//		}
//	}
func myfun(numbers ...interface{}) {
	for _, k := range numbers {
		fmt.Printf("参数值为：%v\n", k)
	}
}

func main() {
	var s = []string{"Mary", "Tim"}
	var m = map[string]interface{}{"name": "Mary", "age": 10}
	// 调用函数
	myfun(45, "Tom", s, m)
}
