package main

import "fmt"

func main() {
	var a, b, c, d = true, true, false, false
	fmt.Printf("a&&b的值为：%v\n", a && b)
	fmt.Printf("a&&c的值为：%v\n", a && c)
	fmt.Printf("c&&d的值为：%v\n", c && d)
	fmt.Printf("a||b的值为：%v\n", a || b)
	fmt.Printf("a||c的值为：%v\n", a || c)
	fmt.Printf("c||d的值为：%v\n", c || d)
	fmt.Printf("!a的值为：%v\n", !a)
	fmt.Printf("!c的值为：%v\n", !c)
}
