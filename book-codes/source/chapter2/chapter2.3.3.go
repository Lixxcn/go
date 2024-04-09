package main

import "fmt"

func main() {
	var a, c int = 21, 0
	c = a
	fmt.Printf("=运算符实例，c值为 = %d\n", c)
	c, a = 1, 20
	c += a
	fmt.Printf("+=运算符实例，c值为 = %d\n", c)
	c, a = 1, 20
	c -= a
	fmt.Printf("-=运算符实例，c值为 = %d\n", c)
	c, a = 1, 20
	c *= a
	fmt.Printf("*=运算符实例，c值为 = %d\n", c)
	c, a = 1, 20
	c /= a
	fmt.Printf("/=运算符实例，c值为 = %d\n", c)
	c, a = 1, 20
	c %= a
	fmt.Printf("求余运算符实例，c值为 = %d\n", c)
	c = 200
	c <<= 2
	fmt.Printf("<<=运算符实例，c值为 = %d\n", c)
	c = 200
	c >>= 2
	fmt.Printf(">>=运算符实例，c值为 = %d\n", c)
	c = 200
	c &= 2
	fmt.Printf("&=运算符实例，c值为 = %d\n", c)
	c = 200
	c ^= 2
	fmt.Printf("^=运算符实例，c值为 = %d\n", c)
	c = 200
	c |= 2
	fmt.Printf("|=运算符实例，c值为 = %d\n", c)
}
