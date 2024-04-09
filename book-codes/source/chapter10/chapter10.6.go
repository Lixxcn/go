package main

import (
	"fmt"
	"reflect"
)

func main() {
	s := []int{1, 2, 3, 4}
	sr := reflect.ValueOf(s)
	sl := sr.Len()
	fmt.Printf("Len()获取切片长度：%v，数据类型：%T\n", sl, sl)
	// 获取切片内存
	srp := sr.Pointer()
	fmt.Printf("Pointer()获取切片内存：%v，数据类型：%T\n", srp, srp)
	si := sr.Index(0)
	fmt.Printf("Index()读取某个元素：%v，数据类型：%T\n", si, si)
	// Set()修改切片元素
	si.Set(reflect.ValueOf(666))
	fmt.Printf("Set()修改某个元素：%v，数据类型：%T\n", sr, sr)
	// Slice3()第一个参数是切片元素的起始索引
	// 第二参数是提取的元素个数，第三个参数是切片元素的终止索引
	s3 := sr.Slice3(0, 3, 4)
	fmt.Printf("Slice3()截取元素：%v，数据类型：%T\n", s3, s3)
	// Slice()第一个参数是切片元素的起始索引
	// 第二个参数是切片元素的终止索引
	ss := sr.Slice(0, 1)
	fmt.Printf("Slice()截取元素：%v，数据类型：%T\n", ss, ss)
	// 反射切片变量转换接口变量，再由接口变量转换切片变量
	srr := sr.Interface().([]int)
	fmt.Printf("反射转换切片：%v\n", srr)
	// 为反射切片变量添加新的元素
	sr = reflect.Append(sr, reflect.ValueOf(666))
	fmt.Printf("Append()添加切片元素：%v\n", sr)
	// 两个反射切片变量合并一个新的反射切片变量
	sr = reflect.AppendSlice(sr, reflect.ValueOf([]int{777}))
	fmt.Printf("AppendSlice()添加合并切片：%v\n", sr)

	// 获取切片的数据类型
	sss := reflect.TypeOf(s)
	// 根据反射切片变量sss创建新的切片
	// MakeSlice()第一个参数是反射切片变量sss
	// 第二个参数切片长度，第三个参数切片容量
	nss := reflect.MakeSlice(sss, 0, 0)
	fmt.Printf("MakeSlice()创建新切片：%v，数据类型：%T\n", nss, nss)
	// 为新切片添加元素
	nss = reflect.Append(nss, reflect.ValueOf(100))
	fmt.Printf("Append()添加新切片元素：%v\n", nss)
}
