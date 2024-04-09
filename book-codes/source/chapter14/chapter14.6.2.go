package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func main() {
	// OpenFile()创建或打开文件，设置读写模式
	// O_RDWR已经支持文件读写操作
	// O_CREATE当文件不存在会自动创建文件
	fs, _ := os.OpenFile("input.csv", os.O_RDWR|os.O_CREATE, 0)
	// 将文件对象fs加载到NewReader()，实例化结构体Reader
	csvReader := csv.NewReader(fs)
	// 一行一行的读取文件，常用于大文件
	for {
		// 调用结构体方法Read()读取文件内容
		row, err := csvReader.Read()
		if err == io.EOF || err != nil {
			break
		}
		fmt.Printf("Read()读取CSV内容：%v，数据类型：%T\n", row, row)
	}
	// 关闭文件
	fs.Close()

	// 一次性读取文件所有内容，常用于小文件
	fs1, _ := os.OpenFile("input.csv", os.O_RDWR|os.O_CREATE, 0)
	// 将文件对象fs1加载到NewReader()，实例化结构体Reader
	csvReader1 := csv.NewReader(fs1)
	// 调用结构体方法ReadAll()读取文件所有内容
	content, err := csvReader1.ReadAll()
	if err != nil {
		fmt.Printf("ReadAll()读取失败：%v\n", err)
	}
	// 遍历输出每一行数据
	for _, row := range content {
		fmt.Printf("ReadAll()读取CSV内容：%v，数据类型：%T\n", row, row)
	}
	// 关闭文件
	fs1.Close()
}
