package main

import "fmt"

func main() {
	for i := 1; i < 5; i++ {
		if i == 2 {
			goto gofunc
		}
		fmt.Printf("本次循环次数为：%v\n", i)
	}
gofunc:
	fmt.Printf("使用goto跳转\n")
	fmt.Printf("程序结束了")
}
