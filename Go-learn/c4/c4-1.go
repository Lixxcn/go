package main

import "fmt"

func main() {
	var a = 65536
	fmt.Println("a value is ", a)
	fmt.Println("a pinter value is ", &a)

	var b = 65535
	p := &b
	fmt.Println("pinter p is ", p)

}
