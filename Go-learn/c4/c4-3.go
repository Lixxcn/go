package main

import "fmt"

func main() {
	s := make([]int, 5, 10)
	fmt.Printf("len %d\n", len(s))
	fmt.Printf("cap %d\n", cap(s))

	for i := 0; i < len(s); i++ {
		s[i] = i
	}
	for _, e := range s {
		fmt.Println(e)
	}

}
