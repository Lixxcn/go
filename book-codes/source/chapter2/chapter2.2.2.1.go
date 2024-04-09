package main

import "fmt"

func main() {

	const (
		_          = iota             // 忽略iota第一个值
		KB float64 = 1 << (10 * iota) // 1 << (10*1)
		MB                            // 1 << (10*2)
		GB                            // 1 << (10*3)
		TB                            // 1 << (10*4)
	)
	fmt.Printf("B转KB的进制为：%.0f\n", KB)
	fmt.Printf("B转MB的进制为：%.0f\n", MB)
	fmt.Printf("B转GB的进制为：%.0f\n", GB)
	fmt.Printf("B转TB的进制为：%.0f\n", TB)
}
