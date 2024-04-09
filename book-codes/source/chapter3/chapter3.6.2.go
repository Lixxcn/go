package main

import "fmt"

func main() {
	// 输出字符a的ASCII
	fmt.Printf("格式化符号v：%v\n", 'a')
	// 输出整型的数据类型
	fmt.Printf("格式化符号T：%T\n", 123)
	// 输出带双引号的字符串
	fmt.Printf("格式化符号q：%q\n", "Hello go")
	// 输出的字符串
	fmt.Printf("格式化符号s：%s\n", "Hello go")
	// 输出保留小数点两位的浮点数，.2是小数点后保留位数
	fmt.Printf("格式化符号f：%.2f\n", 123.321)
	// 输出十进制的整型
	fmt.Printf("格式化符号d：%d\n", 3121)
	//// 输出十进制的整型
	//fmt.Printf("格式化符号d：%d\n", 12.12)
}
