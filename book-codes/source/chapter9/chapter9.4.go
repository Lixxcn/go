package main

import (
	"fmt"
)

// 定义接口
type leg interface {
	leg_run()
}

type mouth interface {
	mouth_speak(content string)
}

// 接口嵌套
type actions interface {
	leg
	mouth
	actions_run()
}

// 定义结构体
type cat struct {
	name string
}

// 定义结构体方法
func (c *cat) mouth_speak(content string) {
	fmt.Printf("%v在说话：%v\n", c.name, content)
}

func (c *cat) leg_run() {
	fmt.Printf("%v在跑步\n", c.name)
}

func (c *cat) actions_run() {
	fmt.Printf("%v在奔跑\n", c.name)
}

// 定义函数
func factory() actions {
	c := cat{name: "凯蒂猫"}
	return &c
}

func main() {
	// 实例化结构体
	c := factory()
	// 调用函数
	c.leg_run()
	c.mouth_speak("喵喵喵喵")
	c.actions_run()
}
