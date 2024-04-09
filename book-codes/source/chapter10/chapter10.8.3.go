package main

import (
	"fmt"
	"reflect"
)

// 定义结构体
type cat struct {
	Name string
	Age  int `json:"age" id:"101"`
}

// 定义结构体方法，指针接收者
func (c *cat) Speak() {
	fmt.Printf("喵...喵...喵\n")
}

// 定义结构体方法，值接收者
func (c cat) Talk() {
	fmt.Printf("喵...喵...喵\n")
}

// 定义结构体方法，带返回值
func (c cat) Sleep() string {
	fmt.Printf("Z...Z...Z\n")
	return "Sleep"
}

// 定义结构体方法，带参数和返回值
func (c cat) Run(a string) {
	fmt.Printf("run...run...%v\n", a)
}

// 定义结构体方法，带参数和返回值
func (c cat) Eat(a string) string {
	fmt.Printf("chi...chi...%v\n", a)
	return "Eat"
}

func main() {
	// 创建结构体变量
	c := cat{Name: "Lily", Age: 18}

	// TypeOf()创建反射结构体变量
	vt := reflect.TypeOf(c)
	// 创建反射结构体指针变量
	//vt := reflect.TypeOf(&c)
	// NumMethod()获取所有结构体方法
	fmt.Printf("NumMethod()获取所有结构体方法：%v\n", vt.NumMethod())
	vmm, _ := vt.MethodByName("Talk")
	fmt.Printf("Func获取方法的内存地址：%v\n", vmm)
	// 遍历输出每个方法的信息
	for i := 0; i < vt.NumMethod(); i++ {
		// 遍历NumMethod()，通过Method(i).Name获取方法名
		fmt.Printf("Name获取方法名：%v\n", vt.Method(i).Name)
		fmt.Printf("PkgPath获取方法所在包名：%v\n", vt.Method(i).PkgPath)
		fmt.Printf("Func获取方法的内存地址：%v\n", vt.Method(i).Func)
		fmt.Printf("Type获取方法的类型：%v\n", vt.Method(i).Type)
		fmt.Printf("Index获取方法的索引：%v\n", vt.Method(i).Index)
	}

	// 创建反射结构体指针变量
	//vc := reflect.ValueOf(&c)
	// 创建反射结构体变量
	vc := reflect.ValueOf(c)
	// 获取结构体方法
	ctn := vc.MethodByName("Eat")
	fmt.Printf("MethodByName()获取方法的内存地址：%v\n", ctn)
	// 获取结构体方法
	cty := vc.Method(0)
	fmt.Printf("Method()获取方法的内存地址：%v\n", cty)
	// 获取方法的类型
	fmt.Printf("Type()获取方法的类型：%v\n", cty.Type())
}
