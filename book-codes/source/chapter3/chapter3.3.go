package main

import "fmt"

func main() {
	var comp1 complex128 = complex(1, 3)
	var comp2 complex128 = complex(2, 4)
	var comp3 complex64 = complex(3, 6)
	adds := comp1 + comp2
	reduce := comp1 - comp2
	mult := comp1 * comp2
	// 不同的复数类型不能直接运算
	// 将complex64转为complex128类型，再执行运算
	div := complex128(comp3) / comp1
	fmt.Printf("复数相加：%v\n", adds)
	fmt.Printf("复数相减：%v\n", reduce)
	fmt.Printf("复数相乘：%v\n", mult)
	fmt.Printf("复数相除：%v\n", div)
}
