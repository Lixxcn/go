package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// 使用os创建或打开文件对象
	f, _ := os.OpenFile("output.txt", os.O_RDWR|os.O_SYNC, 0)
	// 实例化结构体Writer
	reader := bufio.NewWriter(f)
	// 调用结构体方法Write()写入数据
	n1, _ := reader.Write([]byte("6666\n"))
	fmt.Printf("Write()已写入数据：%v\n", n1)
	// 调用结构体方法WriteString()写入数据
	n2, _ := reader.WriteString("7777\n")
	fmt.Printf("WriteString()已写入数据：%v\n", n2)
	// 调用结构体方法Flush()将数据保存到文件中
	reader.Flush()
}
