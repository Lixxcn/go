package main

import "fmt"

// 定义接口
type actions interface {
	// 没有参数没有返回值
	walk()
	// 没有参数有返回值
	runs() (int, int)
	// 有参数没有返回值
	speak(content string, speed int)
	// 有参数有返回值
	rest(sleepTime int) int
}

// 定义结构体
type cats struct {
	name string
}

// 定义接口方法的功能逻辑
func (c *cats) walk() {
	fmt.Printf("%v在散步\n", c.name)
}

func (c *cats) runs() (int, int) {
	fmt.Printf("%v在跑步\n", c.name)
	speed := 10
	time := 1
	return speed, time
}

func (c *cats) speak(content string, speed int) {
	fmt.Printf("%v在说话：%v，语速：%v\n", c.name, content, speed)
}

func (c *cats) rest(sleepTime int) int {
	fmt.Printf("%v在休息，入睡时间：%v小时\n", c.name, sleepTime)
	return sleepTime
}

func main() {
	// 定义接口变量
	var a actions
	// 结构体实例化
	c := cats{name: "kitty"}
	// 结构体实例化变量的指针赋值给接口变量
	a = &c
	// 调用接口里的方法
	a.walk()
	speed, time := a.runs()
	fmt.Printf("跑步速度：%v，跑步时间：%v\n", speed, time)
	a.speak("喵喵", 2)
	sleepTime := a.rest(10)
	fmt.Printf("入睡时间：%v小时\n", sleepTime)
}
