package main

import "fmt"

func main() {
	var ss = []int{1, 2}
	fmt.Printf("新增元素前的切片ss：%v\n", ss)
	// 新增元素不覆盖原有切片
	sss := append(ss, 3)
	fmt.Printf("新增元素后的切片ss：%v\n", ss)
	fmt.Printf("新切片sss：%v\n", sss)
	// 新增元素并覆盖原有切片
	ss = append(ss, 4)
	fmt.Printf("新增元素后的切片ss：%v\n", ss)
	// 添加多个元素
	ss = append(ss, 5, 6, 7, 8)
	fmt.Printf("新增元素后的切片ss：%v\n", ss)
}
