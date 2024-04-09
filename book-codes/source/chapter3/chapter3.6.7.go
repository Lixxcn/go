package main

import (
	"fmt"
	"strings"
)

func main() {
	n := "hello@-world@-I@-am@-Tom"
	// 对字符串的空格进行分割
	m := strings.Split(n, "@-")
	fmt.Printf("分割后的数据类型：%T\n", m)
	for _, i := range m {
		// 输出分割后的每个字符串
		fmt.Println(i)
	}
}
