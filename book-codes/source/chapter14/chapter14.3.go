package main

import (
	"fmt"
	"os"
)

func main() {
	// 定义文本内容
	var val = "This is values\n"
	// 将文本内容转为字节类型的切片
	var val_byte = []byte(val)
	fmt.Printf("values_byte的数据长度：%v\n", len(val_byte))
	// 创建或打开文件
	// 不能使用Open()和Create()，因为文件读写模式不符合要求
	f, _ := os.OpenFile("output.txt", os.O_RDWR|os.O_SYNC|os.O_APPEND, 0)

	// Write()往文件写入文本内容（字节类型的切片）
	n1, _ := f.Write(val_byte)
	fmt.Printf("Write()写入数据的长度：%v\n", n1)

	// WriteString()往文件写入文本内容（字符串类型）
	n2, _ := f.WriteString(val)
	fmt.Printf("WriteString()写入数据的长度：%v\n", n2)

	// WriteAt()相当于Write()+Seek()的功能
	// Seek()是文本内容的偏移量，将写入数据在已有数据的某个位置开始写入
	// WriteAt()不支持O_APPEND模式
	n3, _ := f.WriteAt(val_byte, 0)
	fmt.Printf("WriteAt()写入数据的长度：%v\n", n3)

	// 读取文件，必须为切片定义长度
	var valu_byte []byte = make([]byte, 10)
	// 读取方式一：全内容读取
	//vb, _ := f.Read(valu_byte)
	// 读取方式二：部分内容读取
	vb, _ := f.ReadAt(valu_byte, 0)
	vbs := string(valu_byte)
	fmt.Printf("ReadAt()读取是数据长度：%v\n", vb)
	fmt.Printf("ReadAt()读取是数据：%v\n", vbs)
	// 关闭文件
	f.Close()
}
