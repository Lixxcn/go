package main

import (
	"fmt"
	"reflect"
)

// 定义无参数无返回值的函数
func myfunc() {
	fmt.Printf("This is myfunc, 6666\n")
}

// 定义带参数无返回值的函数
func myfunc1(name string) {
	fmt.Printf("This is myfunc1, para is %v\n", name)
}

// 定义无参数带返回值的函数
func myfunc2() string {
	fmt.Printf("This is myfunc2, 6666\n")
	return "6666"
}

// 定义带参数带返回值的函数
func myfunc3(name string) string {
	fmt.Printf("This is myfunc3, para is %v\n", name)
	return "7777"
}

func main() {
	// 反射无参数无返回值的函数
	mf := reflect.ValueOf(myfunc)
	// 判断反射函数变量的类型
	fmt.Println("rf is reflect.Func?", mf.Kind() == reflect.Func)
	// 调用反射函数，无参数可设为nil
	mf.Call(nil)

	// 反射带参数无返回值的函数
	mf1 := reflect.ValueOf(myfunc1)
	// 判断反射函数变量的类型
	fmt.Println("rf1 is reflect.Func?", mf1.Kind() == reflect.Func)
	// 调用反射函数，函数参数必须为reflect.Value类型的切片
	// 切片元素顺序对应函数参数的顺序
	mf1.Call([]reflect.Value{reflect.ValueOf("Tom")})

	// 反射无参数带返回值的函数
	mf2 := reflect.ValueOf(myfunc2)
	// 判断反射函数变量的类型
	fmt.Println("rf2 is reflect.Func?", mf2.Kind() == reflect.Func)
	// 调用反射函数，无参数可设为nil
	myr2 := mf2.Call(nil)
	// 输出返回值的数据和类型
	fmt.Printf("myfunc2 return is %v, %T\n", myr2, myr2)

	// 反射带参数带返回值的函数
	mf3 := reflect.ValueOf(myfunc3)
	// 判断反射函数变量的类型
	fmt.Println("mf3 is reflect.Func?", mf3.Kind() == reflect.Func)
	// 调用反射函数，函数参数必须为reflect.Value类型的切片
	// 切片元素顺序对应函数参数的顺序
	myr3 := mf3.Call([]reflect.Value{reflect.ValueOf("Tom")})
	// 输出返回值的数据和类型
	fmt.Printf("myfunc2 return is %v, %T\n", myr3, myr3)
}
