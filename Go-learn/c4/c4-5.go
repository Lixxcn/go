package main

import "fmt"

func main() {
	s := make([]int, 5, 10)
	fmt.Printf("pointer %p\n", s)

	s = append(s, 5, 6, 7, 8, 9)
	fmt.Printf("pointer %p\n", s)

	s = append(s, 10)
	fmt.Printf("pointer %p\n", s)

}
