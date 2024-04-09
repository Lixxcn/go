package main

import "fmt"

func main() {
	myStr := []string{"Jack", "Mark"}
	// 使用for循环输出切片myStr的元素
	for i := 0; i < len(myStr); i++ {
		fmt.Printf("本次循环的次数为：%v\n", i)
		fmt.Printf("切片myStr的元素为：%v\n", myStr[i])
	}
	// 使用for-range输出切片myStr的元素
	for i, v := range myStr {
		fmt.Printf("本次循环的次数为：%v\n", i)
		fmt.Printf("切片myStr的元素为：%v\n", v)
	}
}
