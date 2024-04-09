package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var weight int
	fmt.Printf("输入你的体重（kg）：")
	fmt.Scan(&weight)
	fmt.Printf("\n")
	if weight < 40 {
		fmt.Printf("体重值为%v，偏轻\n", weight)
	} else if 40 <= weight && weight <= 70 {
		fmt.Printf("体重值为%v，正常\n", weight)
	} else {
		fmt.Printf("体重值为%v，偏重\n", weight)
	}

	// 随机数
	rand.Seed(time.Now().Unix())
	// num := rand.Intn(100)从100中随机生成整数
	if num := rand.Intn(100); num < 20 {
		fmt.Printf("随机数为%v\n", num)
	} else if num > 20 {
		fmt.Printf("随机数为%v\n", num)
	}

}
