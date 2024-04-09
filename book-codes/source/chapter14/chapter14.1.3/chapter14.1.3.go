package main

import (
	"fmt"
	"github.com/axgle/mahonia"
	"os/exec"
)

func main() {
	// 设置gbk编码格式
	enc := mahonia.NewDecoder("gbk")
	// 执行ipconfig指令
	cmd := exec.Command("cmd", "/C", "ipconfig")
	// 获取指令的输出结果
	output, err := cmd.Output()
	// 若有错误则输出异常信息
	if err != nil {
		fmt.Println(err)
	}
	// 因为结果是字节数组，需要转换成字符串
	// enc.ConvertString是对字符串进行编码处理
	fmt.Println(enc.ConvertString(string(output)))
}
