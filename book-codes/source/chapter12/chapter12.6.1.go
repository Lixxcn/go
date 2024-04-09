package main

import "fmt"

func main() {
	// 为变量myInt创建Int类型的内存地址
	myInt := new(int)
	fmt.Printf("myInt类型：%T，数值：%v，地址：%v\n", *myInt, *myInt, myInt)

	// 为变量myStr创建String类型的内存地址
	myStr := new(string)
	fmt.Printf("myStr类型：%T，数值：%v，地址：%v\n", *myStr, *myStr, myStr)

	// 为变量myMap创建Map类型的内存地址
	myMap := new(map[string]string)
	fmt.Printf("myMap类型：%T，数值：%v，地址：%v\n", *myMap, *myMap, myMap)

	// 为变量mySli创建Slice类型的内存地址
	mySli := new([]int)
	fmt.Printf("mySli类型：%T，数值：%v，地址：%v\n", *mySli, *mySli, mySli)

	// 为变量myChan创建Channel类型的内存地址
	myChan := new(chan string)
	fmt.Printf("myChan类型：%T，数值：%v，地址：%v\n", *myChan, *myChan, myChan)
}
