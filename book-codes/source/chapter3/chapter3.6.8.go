package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "hello world I am Tom"
	// 参数n代表替换的次数，从左边开始计算，-1代表替换全部
	m := strings.Replace(s, " ", "-", -1)
	fmt.Println(m)
	// 参数n等于1代表只替换一次
	k := strings.Replace(s, " ", "-", 1)
	fmt.Println(k)
}
