package main

import "fmt"

func main() {

	var a = 60 // 60 = 0011 1100
	var b = 13 // 13 = 0000 1101
	var c = 0

	c = a & b // 12 = 0000 1100
	fmt.Printf("c的十进制值为 %d\n", c)
	fmt.Printf("c的二进制值为 %b\n", c)

	c = a | b // 61 = 0011 1101
	fmt.Printf("c的十进制值为 %d\n", c)
	fmt.Printf("c的二进制值为 %b\n", c)

	c = a ^ b // 49 = 0011 0001
	fmt.Printf("c的十进制值为 %d\n", c)
	fmt.Printf("c的二进制值为 %b\n", c)

	c = a << 2 // 240 = 1111 0000
	fmt.Printf("c的十进制值为 %d\n", c)
	fmt.Printf("c的二进制值为 %b\n", c)

	c = a >> 2 // 15 = 0000 1111
	fmt.Printf("c的十进制值为 %d\n", c)
	fmt.Printf("c的二进制值为 %b\n", c)
}
