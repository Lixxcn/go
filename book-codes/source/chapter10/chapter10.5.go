package main

import (
	"fmt"
	"reflect"
)

type myint int

type cat struct {
	name string
}

func main() {
	var n int = 55
	rn := reflect.TypeOf(n)
	fmt.Printf("反射变量rn的类型：%v\n", rn)
	fmt.Printf("反射变量rn的类型：%v\n", rn.Name())
	fmt.Printf("反射变量rn所属的种类：%v\n", rn.Kind())

	var x myint = 66
	rx := reflect.TypeOf(x)
	fmt.Printf("反射变量rx的类型：%v\n", rx)
	fmt.Printf("反射变量rx的类型：%v\n", rx.Name())
	fmt.Printf("反射变量rx所属的种类：%v\n", rx.Kind())

	c := cat{name: "Lily"}
	vc := reflect.TypeOf(c)
	fmt.Printf("反射变量vc的类型：%v\n", vc)
	fmt.Printf("反射变量vc的类型：%v\n", vc.Name())
	fmt.Printf("反射变量vc所属的种类：%v\n", vc.Kind())
}
