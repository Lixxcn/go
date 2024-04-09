package main

import "fmt"

func main() {
	n := 123
	f := 123.456
	// 使用int()转为整型
	fmt.Printf("浮点型转整型：%T：%v\n", int(f), int(f))
	// 使用float64()或float32()转为浮点型
	fmt.Printf("整型转浮点型：%T：%v\n", float64(n), float64(n))
	// 整型与浮点型的运算
	m := int(f) + n
	k := float64(n) + f
	fmt.Printf("转整型再运算：%T：%v\n", m, m)
	fmt.Printf("转浮点型再运算：%T：%v\n", k, k)
}
