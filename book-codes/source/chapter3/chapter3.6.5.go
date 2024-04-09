package main

import "fmt"

func main() {
	//n := "Hi,Go语言"
	//for i:=0; i<len(n); i++{
	//	// %c输出每个字符
	//	// %d输出字符对应的十进制数
	//	fmt.Printf("%c——%d\n", n[i], n[i])
	//}

	//n := "Hi,Go语言"
	//m := []rune(n)
	//for i:=0; i<len(m); i++{
	//	// %c输出每个字符
	//	// %d输出字符对应的十进制数
	//	fmt.Printf("%c——%d\n", m[i], m[i])
	//}

	n := "Hi,Go语言"
	for _, s := range n {
		// %c输出每个字符
		// %d输出字符对应的十进制数
		fmt.Printf("%c——%d\n", s, s)
	}

}
