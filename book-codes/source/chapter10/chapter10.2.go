package main

import (
	"fmt"
	"reflect"
)

func main() {
	it := reflect.TypeOf(32)
	iv := reflect.ValueOf(32)
	fmt.Printf("整型类型：%v，反射类型：%T\n", it, it)
	fmt.Printf("整型的数值：%v，反射类型：%T\n", iv, iv)
	istrt := reflect.TypeOf("abc")
	istrv := reflect.ValueOf("abc")
	fmt.Printf("字符串类型：%v，反射类型：%T\n", istrt, istrt)
	fmt.Printf("字符串的数值：%v，反射类型：%T\n", istrv, istrv)
}
