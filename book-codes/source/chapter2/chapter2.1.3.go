package main

// 场合一
// 只使用net/http/pprof的初始化函数init()
import (
	"fmt"
	_ "net/http/pprof"
)

func myfunc() (int, string) {
	// 自定义函数，设置两个返回值
	a := 10
	b := "golang"
	return a, b
}

type Foo interface {
	// 定义接口
	Say()
}
type Dog struct {
	// 定义结构体
	name string
}

func (d Dog) Say() {
	// 结构体实现接口Foo的使用
	fmt.Println(d.name + " say hi")
}

func main() {
	// 场合二
	// 调用函数myfunc()并只获取第一个返回值
	a, _ := myfunc()
	fmt.Printf("只获取函数myfunc的第一个返回值%d：\n", a)

	// 场合三
	// 判断结构体Dog是否实现接口Foo的使用
	// 等同于判定有没有定义func (d Dog) Say(){}
	// 用作类型断言，如果Dog没有实现Foo，则会报编译错误
	var _ Foo = Dog{"black dog"}
}
