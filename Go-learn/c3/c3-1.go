package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a float32

	b := 1.4

	fmt.Println(reflect.TypeOf(a))

	fmt.Println(reflect.TypeOf(b))
}
