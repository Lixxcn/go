package main

import "fmt"

func main() {
	const (
		a = iota // iota设为0，常量a的值为iota的值
		b        // iota累加1，常量b的值为iota的值
		c = 10   // iota累加1，常量c的值为10
		d        // iota累加1，常量d的值为10
		e = iota // iota累加1，常量e的值为iota的值
	)
	fmt.Printf("a的值为：%d\n", a)
	fmt.Printf("b的值为：%d\n", b)
	fmt.Printf("c的值为：%d\n", c)
	fmt.Printf("d的值为：%d\n", d)
	fmt.Printf("e的值为：%d\n", e)
}
