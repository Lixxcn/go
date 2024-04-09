package main

import (
	"fmt"
	"os"
	"time"
)

// 操作目录与文件
func main() {
	// 获取当前目录
	gw, _ := os.Getwd()
	fmt.Printf("获取当前目录：%v\n", gw)
	// 改变当前工作目录
	os.Chdir("D:/")
	gwn, _ := os.Getwd()
	fmt.Printf("改变当前工作目录：%v\n", gwn)
	// 创建文件，由于当前路径改为D盘，因此在D盘创建文件
	f1, _ := os.Create("./1.txt")
	f1.Close()
	// 修改文件权限
	// 第二个参数mode在Windows系统下
	// mode为0200代表所有者可写
	// mode为0400代表只读
	// mode为0600代表读写
	os.Chmod("D:/1.txt", 0400)
	// 修改文件的访问时间和修改时间
	nows := time.Now().Add(time.Hour)
	os.Chtimes("D:/1.txt", nows, nows)
	// 把字符串中带${var}或$var替换成指定指符串
	s := "你好，${1}${2}$3"
	fmt.Printf(os.Expand(s, func(k string) string {
		mapp := map[string]string{
			"1": "我是",
			"2": "go",
			"3": "语言",
		}
		return mapp[k]
	}))
	// 创建目录
	os.Mkdir("D:/abc", os.ModePerm)
	// 创建多级目录
	os.MkdirAll("D:/abc/d/e/f", os.ModePerm)
	// 删除文件或目录
	os.Remove("D:/abc/d/e/f")
	// 删除指定目录下所有文件
	os.RemoveAll("D:/abc")
	// 重命名文件
	os.Rename("D:/1.txt", "D:/1_new.txt")
	// 判断文件相同。Stat()获取文件信息，SameFile()判断文件相同
	f2, _ := os.Create("D:/2.txt")
	fs2, _ := f2.Stat()
	f3, _ := os.Create("D:/3.txt")
	fs3, _ := f3.Stat()
	fmt.Printf("f2和f3是否同一文件：%v\n", os.SameFile(fs2, fs3))
	// 返回临时目录
	fmt.Printf("返回临时目录：%v\n", os.TempDir())
}
