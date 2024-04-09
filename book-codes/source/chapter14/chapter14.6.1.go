package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	// OpenFile()创建或打开文件，设置读写模式
	// 如果设置O_APPEND模式，实现文件续写功能
	// 如果设置O_TRUNC模式，新数据覆盖文件原有数据
	nfs, _ := os.OpenFile("input.csv", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0)
	// 将文件对象nfs加载到NewWriter()，实例化结构体Writer
	csvWriter := csv.NewWriter(nfs)
	// 设置结构体Writer的成员
	// Comma设置每个字段之间的分隔符，默认为逗号
	csvWriter.Comma = ','
	// UseCRLF默认为true，使用\r\n作为换行符
	csvWriter.UseCRLF = true

	// 写入一行数据
	row := []string{"1", "2", "3", "4"}
	err := csvWriter.Write(row)
	if err != nil {
		fmt.Printf("无法写入，错误信息：%v\n", err)
	}

	// 一次性写入多行数据
	var newContent [][]string
	newContent = append(newContent, []string{"11", "12", "13", "14"})
	newContent = append(newContent, []string{"21", "22", "23", "24"})
	csvWriter.WriteAll(newContent)

	// 将数据写入文件
	csvWriter.Flush()
	// 关闭文件
	nfs.Close()
}
