package main

import (
	"fmt"
	"reflect"
)

func main() {
	// 定义集合
	m := make(map[string]string)
	m["name"] = "Tom"
	m["age"] = "100"
	// 转换反射集合变量
	mr := reflect.ValueOf(m)
	// 获取键值对数量
	fmt.Printf("Len()获取键值对数量：%v\n", mr.Len())
	// 获取集合内存地址
	mrt := mr.Pointer()
	fmt.Printf("Pointer()获取集合内存地址：%v，数据类型：%T\n", mrt, mrt)
	// 获取集合所有的键
	mk := mr.MapKeys()
	fmt.Printf("MapKeys()获取键：%v，数据类型：%T\n", mk, mk)
	// 通过集合的键获取对应键值对
	mi := mr.MapIndex(mk[0])
	fmt.Printf("MapIndex()获取键值对：%v，数据类型：%T\n", mi, mi)
	// 获取集合的所有键值对
	iter := mr.MapRange()
	for iter.Next() {
		// 使用方法Next()输出所有键值对
		k := iter.Key()
		v := iter.Value()
		fmt.Printf("MapRange()获取集合的键：%v，集合的值：%v\n", k, v)
	}
	// 将反射集合变量转换接口变量，再由接口变量转换集合变量
	mm := mr.Interface().(map[string]string)
	fmt.Printf("反射转换集合：%v\n", mm["name"])
	// 添加新的键值对
	newKey := reflect.ValueOf("address")
	newValue := reflect.ValueOf("GuangZhou")
	mr.SetMapIndex(newKey, newValue)
	fmt.Printf("SetMapIndex()添加键值对：%v\n", mr)

	// 获取集合的数据类型
	mmm := reflect.TypeOf(m)
	// 根据反射变量mmm创建新的集合
	nmmm := reflect.MakeMap(mmm)
	// 添加新的键值对
	nmmm.SetMapIndex(newKey, newValue)
	fmt.Printf("MakeMap()创建新的集合：%v\n", nmmm)
}
