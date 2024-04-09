package main

import (
	"fmt"
	"reflect"
)

func main() {
	num := 32
	fmt.Printf("转换前的数据：%v，数据类型：%T\n", num, num)
	iv := reflect.ValueOf(num)
	fmt.Printf("接口转换反射：%v，数据类型：%T\n", iv, iv)
	i := iv.Interface()
	fmt.Printf("反射转换接口：%v，数据类型：%T\n", i, i)
}
