package main

import "fmt"

func main() {
	// 定义长度为2的数组并设置每个元素值
	var s = [2]int{100, 200}
	// 输出数组元素
	for i := 0; i < len(s); i++ {
		fmt.Printf("数组s第%v个元素是：%v\n", i+1, s[i])
	}

	// 定义数组并设置每个元素值，数值长度根据元素个数自动设置
	var ss = [...]int{300, 400}
	// 输出数组元素
	for i := 0; i < len(ss); i++ {
		fmt.Printf("数组ss第%v个元素是：%v\n", i+1, ss[i])
	}

	// 定义数组并设置第一个和第四个元素值
	var sss = [...]int{0: 300, 3: 500}
	// 输出数组元素
	for i := 0; i < len(sss); i++ {
		fmt.Printf("数组sss第%v个元素是：%v\n", i+1, sss[i])
	}
}
