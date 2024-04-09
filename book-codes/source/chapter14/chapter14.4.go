package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// 使用os包创建文件
	f, _ := os.OpenFile("output.txt", os.O_RDWR|os.O_SYNC, 0)
	// 返回ReadCloser对象提供close函数
	f1 := ioutil.NopCloser(f)
	defer f1.Close()

	// ReadAll()读取所有数据
	p, _ := ioutil.ReadAll(f)
	// 将数据从字节切片转换字符串
	fmt.Printf("ReadAll()读取所有数据：%v\n", string(p))

	// ReadDir返回目录下所有文件切片
	fileInfo, _ := ioutil.ReadDir("./")
	for _, data := range fileInfo {
		fmt.Printf("ReadDir()的文件信息：%v\n", data.Name())
	}

	// 读取整个文件数据
	data, _ := ioutil.ReadFile("output.txt")
	// 将数据从字节切片转换字符串
	fmt.Printf("ReadFile()读取文件内容：%v\n", string(data))

	// 往文件写入数据，不支持文件内容续写，每次调用会覆盖原有数据
	ioutil.WriteFile("output.txt", []byte("111"), 0655)

	// 当前路径下创建test前缀的临时文件夹，返回文件夹名称
	dir, _ := ioutil.TempDir("./", "test")
	fmt.Printf("TempDir()创建的文件夹：%v\n", dir)

	// 当前路径下创建test前缀的临时文件，返回os.File指针
	fs, _ := ioutil.TempFile("./", "test")
	// 使用os的Write()写入数据
	fs.Write([]byte("222"))
	fs.Close()
}
