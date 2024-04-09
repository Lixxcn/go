package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var num int
	var products = map[int]int{}
	var sample []int
	// 设置随机数的随机种子
	rand.Seed(time.Now().UnixNano())
	// 输出操作提示
	fmt.Printf("请输入检测产品数量：\n")
	// 存储用户输入数据
	fmt.Scanln(&num)
	// 根据产品数量生成产品编号
	// 编号从1开始，以集合存储
	for i := 1; i <= num; i++ {
		products[i] = i
	}
	fmt.Printf("产品编号：%v\n", products)

	// 产品数量的二分一到四分一范围选择样品数量
	runs := rand.Intn(num/2) + (num / 4)
	// 根据样品数量从产品中随机抽取产品
	// 每次循环只抽取一件产品
	for i := 1; i <= runs; i++ {
		// 随机生成产品编号
		n := rand.Intn(num) + 1
		// 判断产品编号是否在产品中
		_, ok := products[n]
		if ok {
			// 若存在则加入切片sample
			sample = append(sample, n)
			// 产品已转为样品，应从products删除
			delete(products, n)
		} else {
			// 产品编号不存在，说明已转为样品
			// 当前循环变量i递减，即当前循环无效
			i--
		}
	}

	// 样品合格数量，用于计算合格率
	var qualified int = 0
	// 遍历所有样品
	for _, k := range sample {
		// 从1到100范围生成随机数
		probability := rand.Intn(100) + 1
		// 随机数大于50则视为产品合格
		if probability > 50 {
			fmt.Printf("产品编号%v检测合格\n", k)
			// 样品合格数量自增加一
			qualified++
		} else {
			// 随机数小于等于50则视为产品不合格
			fmt.Printf("产品编号%v检测不合格\n", k)
		}
	}
	// 计算合格率（百分比表示）：样品合格数量 / 样品总数
	rate := float64(qualified) / float64(len(sample)) * 100
	fmt.Printf("合格率：%.2f%%\n", rate)
}
