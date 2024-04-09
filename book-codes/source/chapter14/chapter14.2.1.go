package main

import (
	"fmt"
	"path"
)

func main() {
	// 判断路径是不是绝对路径
	fmt.Printf("IsAbs函数：%v\n", path.IsAbs("./a/b"))
	// path.IsAbs仅兼容Linux，不兼容Windows，因此输出false
	fmt.Printf("IsAbs函数：%v\n", path.IsAbs("D:a/b"))
	// 路径拼接，连接后自动调用Clean函数
	fmt.Printf("Join函数：%v\n", path.Join("./a", "b/c", "../d/"))
	// 返回路径的最后一个元素
	fmt.Printf("Base函数：%v\n", path.Base("D:/a/b/c"))
	// 如果路径为空字符串或斜杠，返回实心点或斜杠
	fmt.Printf("Base函数：%v\n", path.Base(""))
	fmt.Printf("Base函数：%v\n", path.Base("/"))
	// 返回等价的最短路径
	// 1.用一个斜线替换多个斜线
	// 2.清除当前路径的实心点
	// 3.清除..和它前面的元素
	// 4.以/..开头的，变成/
	fmt.Printf("Clean函数：%v\n", path.Clean("./a/b/c/../"))
	// 返回路径最后一个元素的前面目录
	// 路径为空则返回实心点
	fmt.Printf("Dir函数：%v\n", path.Dir("D:/a/b/c"))
	// 返回路径中的文件扩展名
	// 如果没有文件扩展名则返回空
	fmt.Printf("Ext函数：%v\n", path.Ext("D:/a/b/c/d.jpg"))
	// 匹配路径，完全匹配则返回true
	// *匹配0或多个非/的字符
	matched1, _ := path.Match("*", "abc")
	fmt.Printf("Match函数：%v\n", matched1)
	// ?匹配一个非/的字符，a?b的?不是斜杠/并只有一个字符都能匹配
	matched2, _ := path.Match("a?b", "agb")
	fmt.Printf("Match函数：%v\n", matched2)
	// 匹配路径是否符合a/*/c格式，如a/abc/c和a/bbb/都能匹配
	matched3, _ := path.Match("a/*/c", "a/bb/c")
	fmt.Printf("Match函数：%v\n", matched3)
	// 分割路径中的目录与文件
	dir, file := path.Split("./a/b/c/d.jpg")
	fmt.Printf("Split函数：目录：%v、文件：%v\n", dir, file)
}
