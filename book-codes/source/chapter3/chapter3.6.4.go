package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	n := "Hello,golang"
	m := "你好,Go语言"
	//fmt.Println("字符串n的长度：", len(n))
	//fmt.Println("字符串m的长度：", len(m))
	// 使用utf8.RuneCountInString()获取字符串长度
	fmt.Println("utf8获取n的长度：", utf8.RuneCountInString(n))
	fmt.Println("utf8获取m的长度：", utf8.RuneCountInString(m))
	// 使用[]rune()获取字符串长度
	fmt.Println("[]rune()获取n的长度：", len([]rune(n)))
	fmt.Println("[]rune()获取m的长度：", len([]rune(m)))
}
