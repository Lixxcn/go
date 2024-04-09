package main

import (
	"fmt"
	"strings"
)

func main() {
	n := "hello-world-world-你好呀"
	// 获取字符串的子串world的最开始位置
	m := strings.Index(n, "world")
	fmt.Println("获取子串world的最开始位置：", m)
	// 获取字符串的子串world在最末端的位置
	l := strings.LastIndex(n, "world")
	fmt.Println("获取world在最末端的位置：", l)
	// 截取m往后的字符串
	k := n[m:]
	fmt.Println("截取m往后的字符串：", k)
	// 截取m位置往后的3位字符串
	p := n[m : m+3]
	fmt.Println("截取m位置往后的3位字符串：", p)
}
