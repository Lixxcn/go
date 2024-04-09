package main

import "fmt"

func main() {
	var ss = []int{1, 2, 3, 4, 5, 6, 7}
	// 截取第二个到第五个元素
	s1 := ss[1:4]
	fmt.Printf("截取第二个到第五个元素：%v\n", s1)
	// 截取第三个元素之后的所有元素
	s2 := ss[2:]
	fmt.Printf("截取第三个元素之后的所有元素：%v\n", s2)
	// 截取第三个元素之前的所有元素
	s3 := ss[:2]
	fmt.Printf("截取第三个元素之前的所有元素：%v\n", s3)
	// 如果切片ss没被覆盖，经过截取后不改变原有的切片数据
	fmt.Printf("切片变量ss的值：%v\n", ss)

	//var ss = []int{1, 2, 3, 4, 5, 6, 7}
	//fmt.Printf("切片ss的元素：%v\n", ss)
	//// 删除元素4、5、6
	//ss = append(ss[:2], ss[6:]...)
	//fmt.Printf("切片ss的元素：%v\n", ss)

}
