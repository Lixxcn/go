package main

import "fmt"

func main() {
	for i := 1; i < 5; i++ {
		if i == 2 {
			continue
		}
		fmt.Printf("本次循环次数为：%v\n", i)
	}
}
