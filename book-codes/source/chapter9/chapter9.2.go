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

// 定义函数
func speaking(a actions, content string) {
	a.speak(content)
}

func main() {
	// 实例化结构体
	d := duck{name: "唐老鸭"}
	c := cat{name: "凯蒂猫"}
	// 调用函数
	speaking(&d, "嘎嘎")
	speaking(&c, "喵喵")
}
