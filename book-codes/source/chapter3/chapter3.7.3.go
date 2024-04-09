package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 浮点型转换字符串
	f := 100.12345678901234567890123456789
	fmt.Println(strconv.FormatFloat(f, 'b', 5, 32))
	fmt.Println(strconv.FormatFloat(f, 'b', 5, 32))
	fmt.Println(strconv.FormatFloat(f, 'e', 5, 32))
	fmt.Println(strconv.FormatFloat(f, 'E', 5, 32))
	fmt.Println(strconv.FormatFloat(f, 'f', 5, 32))
	fmt.Println(strconv.FormatFloat(f, 'g', 5, 32))
	fmt.Println(strconv.FormatFloat(f, 'G', 5, 32))

	// 字符串转换浮点型，类型为float32
	s := "0.1234"
	k, _ := strconv.ParseFloat(s, 32)
	fmt.Printf("字符串转换浮点型，类型为float32：%T——%v\n", k, k)
	// 字符串转换浮点型，类型为float64
	l, _ := strconv.ParseFloat(s, 64)
	fmt.Printf("字符串转换浮点型，类型为float64：%T——%v\n", l, l)
}
