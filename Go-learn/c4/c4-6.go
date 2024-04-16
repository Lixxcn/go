package main

import "fmt"

func main() {
	s := make([]int, 5, 10)
	fmt.Printf("len %d, cap %d, pointer %p.\n", len(s), cap(s), s)

	_ = append(s, 5, 6, 7, 8, 9)
	fmt.Printf("len %d, cap %d, pointer %p.\n", len(s), cap(s), s)

	s1 := append(s, 10)
	fmt.Printf("len %d, cap %d, pointer %p.\n", len(s), cap(s), s)
	fmt.Printf("len %d, cap %d, pointer %p.\n", len(s1), cap(s1), s1)

}
