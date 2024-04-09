package main

import (
	"fmt"
)

func main() {
	//fmt.Println(math.MaxFloat32)
	//fmt.Println(math.MaxFloat64)

	var f32 float32 = 1.1
	var f64 float64 = 2.2
	// 将float32的浮点数转为float64，再执行相加运算
	r := float64(f32) + f64
	fmt.Printf("运行结果为：%v\n", r)
	// 将整型转为float32的浮点数，再执行相加运算
	var i int = 10
	rd := float32(i) + f32
	fmt.Printf("运行结果为：%v\n", rd)

	k := fmt.Sprintf("%.1f", r)
	fmt.Printf("进度丢失处理结果为：%v", k)
}
