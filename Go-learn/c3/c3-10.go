package main

import "fmt"

func main() {
	var a = [5]int{0, 1, 2, 3, 4}
	fmt.Println(a)
	modify(&a)
	fmt.Println(a)

}

func modify(a *[5]int) {
	a[1] = 100
	fmt.Println(*a)
}
