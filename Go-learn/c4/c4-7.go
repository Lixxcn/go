package main

import "fmt"

func main() {
	a := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("a  len %d, cap %d, pointer %p. %v\n", len(a), cap(a), &a, a)

	s := a[:]
	fmt.Printf("s  len %d, cap %d, pointer %p. %v\n", len(s), cap(s), s, s)

	s1 := a[1:3]
	fmt.Printf("s1 len %d, cap %d, pointer %p. %v\n", len(s1), cap(s1), s1, s1)

}
