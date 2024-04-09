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
	// 创建反射结构体指针变量
	vc := reflect.ValueOf(&c)
	// 创建反射结构体变量
	//vc := reflect.ValueOf(c)

	// MethodByName()获取Speak()，Call()调用Speak()
	cs := vc.MethodByName("Speak")
	cs.Call(make([]reflect.Value, 0))

	// MethodByName()获取Talk()，Call()调用Talk()
	ct := vc.MethodByName("Talk")
	ct.Call(make([]reflect.Value, 0))

	// MethodByName()获取Sleep()，Call()调用Sleep()
	css := vc.MethodByName("Sleep")
	r := css.Call(make([]reflect.Value, 0))
	fmt.Printf("Sleep()返回值：%v\n", r)

	// MethodByName()获取Run()，Call()调用Run()
	cr := vc.MethodByName("Run")
	cr.Call([]reflect.Value{reflect.ValueOf("GOGOGO")})

	// MethodByName()获取Eat()，Call()调用Eat()
	cea := vc.MethodByName("Eat")
	rr := cea.Call([]reflect.Value{reflect.ValueOf("mouse")})
	fmt.Printf("Eat()返回值：%v\n", rr)
}
