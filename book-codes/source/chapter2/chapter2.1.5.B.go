package main

import (
	// 导入自定义包aa
	"./aa"
	"fmt"
)

func main() {
	// 调用自定义包的A_string
	fmt.Printf("调用自定义包的A_string：%s", aa.A_string)
	fmt.Printf("2: %d", aa.Number)
	fmt.Println("number", aa.Number)
}
