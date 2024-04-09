package main

import (
	"fmt"
	"path"
)

func makeSuffixFunc(suffix string) func(string) string {
	fmt.Printf("参数suffix的内存地址：%v\n", &suffix)
	// 参数suffix是文件后缀名
	return func(name string) string {
		fmt.Printf("参数suffix在匿名函数的内存地址：%v\n", &suffix)
		// 匿名函数使用makeSuffixFunc的参数suffix
		// 参数name是文件名，可能含有后缀名或没有后缀名
		// path.Ext()获取文件后缀名，判断path.Ext()是否为空
		// 判断结果为空，说明参数name没有后缀名，根据suffix创建后缀名
		// 判断结果不为空，说明参数name已有后缀名
		if path.Ext(name) == "" {
			return name + suffix
		} else {
			return "文件已有后缀名：" + name
		}
	}
}

func main() {
	// 定义jpg文件类型的函数变量
	jpgFunc := makeSuffixFunc(".jpg")
	// 定义txt文件类型的函数变量
	txtFunc := makeSuffixFunc(".txt")
	// 判断文件是否已有后缀名
	// 没有后缀名则根据函数变量自动创建
	fmt.Println(jpgFunc("test.png"))
	fmt.Println(txtFunc("test"))
}
