package main

import (
	"fmt"
	"math"
)

func main() {
	//var result int = 0
	//// 定义3行2列长度的二维数组
	//var s [3][2]int
	//// 为二维数组赋值
	//s = [3][2]int{{10, 20}, {30, 40}, {50, 60}}
	//for i := 0; i < len(s); i++ {
	//	// 循环每一行数据
	//	for k := 0; k < len(s[i]); k++ {
	//		// 循环每一列数据
	//		result = result + s[i][k]
	//		fmt.Printf("当前元素值为：%v\n", s[i][k])
	//	}
	//}
	//fmt.Printf("二维数组的总行数为：%v\n", len(s))
	//fmt.Printf("二维数组的总列数为：%v\n", len(s[0]))
	//fmt.Printf("二维数组的总值为：%v\n", result)

	// 定义2*1*3长度的三维数组
	var point [2][1][3]int
	point = [2][1][3]int{{{3, 5, 7}}, {{5, 3, 2}}}
	// 获取坐标点
	pointA := point[0][0]
	pointB := point[1][0]
	fmt.Printf("坐标点A：%v\n", pointA)
	fmt.Printf("坐标点B：%v\n", pointB)
	// 计算两个坐标点距离
	// 计算两坐标的x坐标之差的平方
	x := (pointA[0] - pointB[0]) * (pointA[0] - pointB[0])
	// 计算两坐标的y坐标之差的平方
	y := (pointA[1] - pointB[1]) * (pointA[1] - pointB[1])
	// 计算两坐标的z坐标之差的平方
	z := (pointA[2] - pointB[2]) * (pointA[2] - pointB[2])
	result := math.Sqrt(float64(x + y + z))
	fmt.Printf("两坐标点距离为：%v\n", result)
}
