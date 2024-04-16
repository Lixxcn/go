package main

import "fmt"

func main() {
	s := make([]int, 5, 10)
	fmt.Printf("len %d\n", len(s))
	fmt.Printf("cap %d\n", cap(s))
	fmt.Printf("pointer %p\n", s)

	s = append(s, 5, 6, 7, 8, 9)
	fmt.Println("=============")
	fmt.Printf("len %d\n", len(s))
	fmt.Printf("cap %d\n", cap(s))
	fmt.Printf("pointer %p\n", s)

}
