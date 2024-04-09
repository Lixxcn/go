package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	//var weight int
	//fmt.Printf("输入你的体重（kg）：")
	//fmt.Scan(&weight)
	//fmt.Printf("\n")
	//if weight < 40 {
	//	fmt.Printf("体重值为%v，偏轻\n", weight)
	//}else if 40 <= weight && weight <= 70{
	//	fmt.Printf("体重值为%v，正常\n", weight)
	//}else {
	//	fmt.Printf("体重值为%v，偏重\n", weight)
	//}

	//var weight, age int
	//fmt.Printf("输入你的年龄：")
	//fmt.Scan(&age)
	//fmt.Printf("\n")
	//fmt.Printf("输入你的体重（kg）：")
	//fmt.Scan(&weight)
	//fmt.Printf("\n")
	//if age < 10 {
	//	fmt.Printf("你的年龄为%v\n", age)
	//	if weight < 15 {
	//		fmt.Printf("体重值为%v，偏轻\n", weight)
	//	}else if weight >= 15 && weight < 30 {
	//		fmt.Printf("体重值为%v，正常\n", weight)
	//	}else{
	//		fmt.Printf("体重值为%v，偏重\n", weight)
	//	}
	//}else if 15 <= age && age <= 30{
	//	fmt.Printf("你的年龄为%v\n", age)
	//	if weight < 40 {
	//		fmt.Printf("体重值为%v，偏轻\n", weight)
	//	}else if weight >= 40 && weight < 60 {
	//		fmt.Printf("体重值为%v，正常\n", weight)
	//	}else{
	//		fmt.Printf("体重值为%v，偏重\n", weight)
	//	}
	//}else {
	//	fmt.Printf("你的年龄为%v\n", age)
	//	if weight < 60 {
	//		fmt.Printf("体重值为%v，偏轻\n", weight)
	//	}else if weight >= 60 && weight < 80 {
	//		fmt.Printf("体重值为%v，正常\n", weight)
	//	}else{
	//		fmt.Printf("体重值为%v，偏重\n", weight)
	//	}
	//}

	// 随机数
	rand.Seed(time.Now().Unix())
	if num := rand.Intn(100); num < 20 {
		fmt.Printf("随机数为%v\n", num)
	} else if num > 20 {
		fmt.Printf("随机数为%v\n", num)
	}
}
