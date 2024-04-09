package main

import "fmt"

// 定义接口
type actions interface {
	speak(content string)
}

// 定义结构体
type duck struct {
	name string
}

type cat struct {
	name string
}

// 定义结构体方法
func (d *duck) speak(content string) {
	fmt.Printf("%v在说话：%v\n", d.name, content)
}

func (c *cat) speak(content string) {
	fmt.Printf("%v在说话：%v\n", c.name, content)
}

// 定义工厂函数
func factory(name string) actions {
	switch name {
	case "duck":
		// 返回结构体duck实例化的内存地址
		return &duck{name: "唐老鸭"}
	case "cat":
		// 返回结构体cat实例化的内存地址
		return &cat{name: "凯蒂猫"}
	default:
		// 自主抛出异常
		panic("No such animal")
	}
}

func main() {
	// 调用工厂函数
	f1 := factory("duck")
	// 调用接口方法speak()
	f1.speak("嘎嘎嘎")
	// 调用工厂函数
	f2 := factory("cat")
	// 调用接口方法speak()
	f2.speak("喵喵喵")
}
