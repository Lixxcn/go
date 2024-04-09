package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 3}
	slice2 := []int{4, 5, 6}
	// 将slice1的元素复制到slice2
	//copy(slice2, slice1)
	// 将slice2的元素复制到slice1
	copy(slice1, slice2)
	fmt.Printf("将slice2的元素复制到slice1：%v\n", slice1)
	fmt.Printf("将slice2的元素复制到slice1：%v\n", slice2)
	slice3 := []int{7, 8, 9, 10}
	// 将slice3的元素复制到slice1
	//copy(slice1, slice3)
	//fmt.Printf("将slice3的元素复制到slice1：%v\n", slice1)
	//fmt.Printf("将slice3的元素复制到slice1：%v\n", slice3)
	// 将slice1的元素复制到slice3
	copy(slice3, slice1)
	fmt.Printf("将slice1的元素复制到slice3：%v\n", slice1)
	fmt.Printf("将slice1的元素复制到slice3：%v\n", slice3)
}
