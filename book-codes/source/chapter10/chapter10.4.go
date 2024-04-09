package main

import (
	"fmt"
	"reflect"
)

func main() {
	//var x int = 66
	//v := reflect.ValueOf(x)
	//v.SetInt(55)

	//var x int = 66
	//v := reflect.ValueOf(x)
	//fmt.Printf("反射变量v能否被修改：%v", v.CanSet())

	var x int = 66
	// reflect.ValueOf的参数设为变量x的指针
	v := reflect.ValueOf(&x)
	fmt.Printf("反射变量v的值：%v\n", v)
	// 通过反射变量v的指针获取变量值
	vv := v.Elem()
	fmt.Printf("反射变量vv的值：%v\n", vv)
	fmt.Printf("反射变量vv能否被修改：%v\n", vv.CanSet())
	vv.SetInt(55)
	fmt.Printf("反射变量vv修改后的值：%v\n", vv)
	fmt.Printf("整型变量x的值：%v\n", x)
}
