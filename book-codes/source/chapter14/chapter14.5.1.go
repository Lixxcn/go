package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// 使用os创建或打开文件对象
	f, _ := os.OpenFile("output.txt", os.O_RDWR|os.O_SYNC, 0)
	reader := bufio.NewReader(f)
	// 通过死循环方式读取每行数据
	for {
		// ReadString()按行读取数据
		buf, err := reader.ReadString('\n')
		// 输出当前读到数据的长度
		fmt.Printf("当前数据长度：%v\n", len(buf))
		// 数据转换字符串格式，并去掉换行符
		value := strings.Trim(string(buf), "\r\n")
		// 如果数据不为空，说明已读到数据
		if value != "" {
			fmt.Printf("当前数据：%v\n", value)
		} else {
			// 若数据为空，可能是空行数据，但出现err则说明已读取完成
			if err != nil {
				break
			} else {
				fmt.Printf("当前数据是空行数据\n")
			}
		}
	}
}
