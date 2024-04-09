package main

import (
	"fmt"
	"reflect"
	"strings"
)

// 定义结构体
type cat struct {
	Name string
	Age  int `json:"age" id:"101"`
}

// 定义结构体方法
func (c *cat) speak() {
	fmt.Printf("喵喵喵")
}

// 定义函数
func GetFieldByIndex(a string) bool {
	return strings.ToLower(a) == "name"
}

func main() {
	// 创建结构体变量
	c := cat{Name: "Lily", Age: 18}
	// ValueOf()创建反射结构体变量
	vc := reflect.ValueOf(c)
	// 计算成员数量
	vnu := vc.NumField()
	fmt.Printf("NumField()计算成员数量：%v，数据类型：%T\n", vnu, vnu)
	// 以成员名称访问成员值
	vn := vc.FieldByName("Name")
	fmt.Printf("FieldByName()访问某个成员：%v，数据类型：%T\n", vn, vn)
	// 以成员排序索引访问成员值
	vi := vc.Field(1)
	fmt.Printf("Field()访问某个成员：%v，数据类型：%T\n", vi, vi)
	// 以成员排序索引访问成员值，索引值以切片表示
	vbi := vc.FieldByIndex([]int{0})
	fmt.Printf("FieldByIndex()访问某个成员：%v，数据类型：%T\n", vbi, vbi)
	// 以函数方式判断并获取某个成员名称，再从成员名称获取成员值
	vf := vc.FieldByNameFunc(GetFieldByIndex)
	fmt.Printf("FieldByNameFunc()访问某个成员：%v，数据类型：%T\n", vf, vf)
	// 判断反射结构体变量能否修改数据
	fmt.Printf("反射结构体变量能否修改数据：%v\n", vc.CanSet())

	// ValueOf()创建反射结构体指针变量
	vc_pit := reflect.ValueOf(&c)
	// 获取所有成员的值
	ve := vc_pit.Elem()
	fmt.Printf("Elem()获取所有成员的值：%v，数据类型：%T\n", ve, ve)
	// 判断反射结构体指针变量能否修改数据
	fmt.Printf("反射结构体指针变量能否修改数据：%v\n", ve.CanSet())
	// Set()、SetInt()等方法设置成员值
	ve.FieldByName("Name").SetString("Tom")
	ve.FieldByName("Name").Set(reflect.ValueOf("Tim"))
	ve.FieldByName("Age").SetInt(666)
	fmt.Printf("Set()、SetInt()等方法设置成员值：%v\n", ve)

	// TypeOf()创建反射结构体变量
	vt := reflect.TypeOf(c)
	// 遍历结构体所有成员数量
	for i := 0; i < vt.NumField(); i++ {
		// 获取每个成员的结构体成员类型
		vinfo := vt.Field(i)
		// 输出成员名和tag
		fmt.Printf("结构成员：%v，其标签为：%v\n", vinfo.Name, vinfo.Tag)
	}
	// 通过成员名, 找到成员类型信息
	if cn, ok := vt.FieldByName("age"); ok {
		// 从tag中取出需要的tag
		fmt.Printf("标签json的内容：%v\n", cn.Tag.Get("json"))
		fmt.Printf("标签id的内容：%v\n", cn.Tag.Get("id"))
	}
	// 通过成员索引, 找到成员类型信息，索引以切片形式表示
	ct := vt.FieldByIndex([]int{1})
	fmt.Printf("标签json的内容：%v\n", ct.Tag.Get("json"))
	fmt.Printf("标签id的内容：%v\n", ct.Tag.Get("id"))
}
