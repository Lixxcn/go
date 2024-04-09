package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 购买号码
	// 定义变量myNum，存放用户当前输入数据
	var myNum int
	// 定义变量myNums，存放用户所有输入数据
	var myNums []int
	// 循环7次，给用户输入7个数据
	for i := 0; i < 7; i++ {
		// 输出操作提示
		fmt.Printf("请输入第%v位号码：\n", i+1)
		// 存储用户输入数据
		fmt.Scanln(&myNum)
		// 将当前数据存放在切片myNums
		myNums = append(myNums, myNum)
	}
	fmt.Printf("你选到号码分别为：%v\n", myNums)

	// 公布开奖号码
	// 定义变量s，切片类型，切片元素为指针类型
	var result []*int
	// 定义变量status，数据类型为布尔型
	var status bool
	// 设置随机数的随机种子
	rand.Seed(time.Now().UnixNano())
	// 设置死循环
	for {
		// 定义变量num，数据类型为整型
		var num int
		// 设置变量status的值
		status = false
		// 创建随机数
		num = rand.Intn(36) + 1
		// 遍历切片result的每个元素
		// 如果随机数num已存在切片result，将变量status等于true
		for _, k := range result {
			if *k == num {
				status = true
			}
		}
		// 变量status等于false
		// 说明随机数num不在切片result里面，将随机数num加入切片result
		if status == false {
			result = append(result, &num)
		}
		// 切片长度等于7，终止死循环
		if len(result) == 7 {
			break
		}
	}
	// 遍历输出切片所有元素
	for i, k := range result {
		fmt.Printf("第%v位号码为：%v\n", i+1, *k)
	}

	// 兑奖
	// 遍历切片result和myNums，将两个切片元素一一对比
	for _, k := range result {
		for _, j := range myNums {
			if *k == j {
				fmt.Printf("号码%v选中了\n", j)
			}
		}
	}
}
