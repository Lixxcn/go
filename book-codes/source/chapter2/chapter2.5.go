package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 定义变量
	var name string
	var weight, height float64
	// 输入名字
	fmt.Printf("请输入你的名字：\n")
	fmt.Scanln(&name)
	// 输入体重
	fmt.Printf("请输入你的体重kg：\n")
	fmt.Scanln(&weight)
	// 输入身高
	fmt.Printf("请输入你的身高cm：\n")
	fmt.Scanln(&height)
	// 输出用户输入信息
	fmt.Printf("%v的体重kg:%v，身高cm:%v\n", name, weight, height)
	// 单位换算，将cm换成m
	height /= 100
	// 计算BMI
	result := weight / (height * height)
	// 将BMI保留两位小数
	BMI, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", result), 64)
	// 根据BMI判断身体状态
	if BMI < 18.5 {
		fmt.Printf("你的BMI为%v，体重过轻", BMI)
	} else if 18.5 <= BMI && BMI < 24.0 {
		fmt.Printf("你的BMI为%v，体重正常", BMI)
	} else if 24 <= BMI && BMI < 27 {
		fmt.Printf("你的BMI为%v，体重过重", BMI)
	} else if 27 <= BMI && BMI < 30 {
		fmt.Printf("你的BMI为%v，体重轻度肥胖", BMI)
	} else if 30 <= BMI && BMI < 35 {
		fmt.Printf("你的BMI为%v，体重中度肥胖", BMI)
	} else {
		fmt.Printf("你的BMI为%v，体重重度肥胖", BMI)
	}
}
