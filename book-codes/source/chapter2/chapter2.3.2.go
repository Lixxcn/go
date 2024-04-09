package main

import "fmt"

func main() {
	var x, y = 8, 5
	fmt.Printf("x大于y：%v\n", x > y)
	fmt.Printf("x大于或等于y：%v\n", x >= y)
	fmt.Printf("x小于y：%v\n", x < y)
	fmt.Printf("x小于或等于y：%v\n", x <= y)
	fmt.Printf("x等于y：%v\n", x == y)
	fmt.Printf("x不等于y：%v\n", x != y)
}
