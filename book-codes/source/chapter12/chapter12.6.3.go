package main

import "fmt"

func main() {
	// 创建切片类型的变量mySli，切片长度和容量为10
	mySli := make([]int, 10)
	// 对切片第一个元素赋值
	mySli[0] = 666
	// 输出切片信息
	fmt.Printf("mySli数值：%v，长度：%v，容量：%v\n", mySli, len(mySli), cap(mySli))

	// 创建集合类型的变量myMap
	myMap := make(map[string]string)
	// 设置集合的成员值
	myMap["name"] = "Tom"
	// 输出集合信息
	fmt.Printf("myMap数值：%v，成员数量：%v\n", myMap, len(myMap))

	// 创建通道类型的变量myChan，容量为10
	myCh := make(chan string, 10)
	// 往通道写入数据
	myCh <- "hello"
	// 输出通道信息
	fmt.Printf("myCh数值：%v，已用缓存：%v，容量：%v\n", myCh, len(myCh), cap(myCh))
}
