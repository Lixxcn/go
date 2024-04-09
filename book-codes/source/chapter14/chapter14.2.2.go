package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	// 判断路径是不是绝对路径，filepath兼容所有操作系统
	fmt.Printf("IsAbs函数：%v\n", filepath.IsAbs("./a/b/c"))
	fmt.Printf("IsAbs函数：%v\n", filepath.IsAbs("C:/a/b/c"))
	// 返回所给路径的绝对路径
	path, _ := filepath.Abs("go.mod")
	fmt.Printf("Abs函数：%v\n", path)
	// 返回路径最后一个元素
	fmt.Printf("Base函数：%v\n", filepath.Base("D:/a/c/1.txt"))
	// 如果路径为空字符串，返回实心点
	fmt.Printf("Base函数：%v\n", filepath.Base(""))
	// 如果路径有一个或多个斜杠/，返回单个斜杠/
	fmt.Printf("Base函数：%v\n", filepath.Base("//"))
	// 返回等价的最短路径
	// 1.用一个斜线替换多个斜线
	// 2.清除当前路径的实心点
	// 3.清除..和它前面的元素
	// 4.以/..开头的，变成/
	fmt.Printf("Clean函数：%v\n", filepath.Clean("D:/a/b/../c"))
	fmt.Printf("Clean函数：%v\n", filepath.Clean("D:/a/b/../c/1.txt"))
	// 返回路径最后一个元素的前面目录
	// 路径为空则返回实心点
	fmt.Printf("Dir函数：%v\n", filepath.Dir("./a/b/c"))
	fmt.Printf("Dir函数：%v\n", filepath.Dir("D:/a/b/c"))
	// 返回软链接的实际路径
	path2, _ := filepath.EvalSymlinks("go.mod")
	fmt.Printf("EvalSymlinks函数：%v\n", path2)
	// 返回文件路径的扩展名
	// 如果不是文件路径，返回空字符串
	fmt.Printf("Ext函数：%v\n", filepath.Ext("./a/b/c/d.jpg"))
	// 将路径中的/替换为路径分隔符
	fmt.Printf("FromSlash函数：%v\n", filepath.FromSlash("./a/b/c"))
	// 返回路径中所有匹配的文件
	match, _ := filepath.Glob("./*.go")
	fmt.Printf("Glob函数：%v\n", match)
	// 路径拼接，连接后自动调用Clean函数
	fmt.Printf("Join函数：%v\n", filepath.Join("C:/a", "/b", "/c"))
	// *匹配0或多个非/的字符
	matched1, _ := filepath.Match("*", "abc")
	fmt.Printf("Match函数：%v\n", matched1)
	// ?匹配一个非/的字符，a?b的?不是斜杠/并只有一个字符都能匹配
	matched2, _ := filepath.Match("a?b", "agb")
	fmt.Printf("Match函数：%v\n", matched2)
	// 匹配路径是否符合a/*/c格式，如a/abc/c和a/bbb/都能匹配
	matched3, _ := filepath.Match("a/*/c", "a/bb/c")
	fmt.Printf("Match函数：%v\n", matched3)
	// 返回参数basepath为基准的相对路径
	path3, _ := filepath.Rel("C:/a/b", "C:/a/b/../e")
	fmt.Printf("Rel函数：%v\n", path3)
	// 将路径使用路径列表分隔符分开，见os.PathListSeparator
	// linux默认为冒号，windows默认为分号
	sl := filepath.SplitList("C:/windows;C:/windows/system")
	fmt.Printf("SplitList函数：%v，长度：%v\n", sl, len(sl))
	// 分割路径中的目录与文件
	dir, file := filepath.Split("C:/a/d.jpg")
	fmt.Printf("Split函数：目录：%v，文件：%v\n", dir, file)
	// 将路径分隔符使用/替换
	fmt.Printf("ToSlash函数：%v\n", filepath.ToSlash("C:\\a\\b"))
	// 返回分区名
	vn := filepath.VolumeName("C:/a/b/c")
	fmt.Printf("VolumeName函数：%v\n", vn)
	// 遍历指定目录下所有文件
	filepath.Walk("./", func(path string, info os.FileInfo, err error) error {
		fmt.Printf("Walk函数：%v\n", path)
		return nil
	})
}
