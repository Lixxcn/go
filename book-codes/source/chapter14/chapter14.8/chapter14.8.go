package main

import (
	"fmt"
	"github.com/xuri/excelize/v2"
)

func main() {
	// NewFile()创建新的Excel文件
	f := excelize.NewFile()
	// NewSheet()在Excel里面创建Sheet2
	index := f.NewSheet("Sheet2")
	// 在Sheet2的单元格写入数据
	f.SetCellValue("Sheet2", "A2", "Hello world")
	f.SetCellValue("Sheet1", "B2", 100)
	// 设置工作簿的默认工作表
	f.SetActiveSheet(index)
	// 从Sheet2获取单元格A2的值
	cell, _ := f.GetCellValue("Sheet2", "A2")
	fmt.Printf("Sheet2的单元格A2的值：%v\n", cell)
	// 保存Excel文件
	err := f.SaveAs("./Book1.xlsx")
	if err != nil {
		fmt.Println(err)
	}
}
