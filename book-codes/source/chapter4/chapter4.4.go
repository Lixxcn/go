package main

import "fmt"

func main() {
	// 组合方式一
	for i := 1; i < 10; i++ {
		fmt.Printf("本次循环：%v\n", i)
	}
	// 组合方式二
	for i := 1; i < 10; {
		fmt.Printf("本次循环：%v\n", i)
		i++
	}
	// 组合方式三
	var i int = 1
	for i < 5 {
		fmt.Printf("本次循环：%v\n", i)
		i++
	}
	// 组合方式四
	for {
		fmt.Printf("本次循环：%v\n", i)
		break
	}
}
