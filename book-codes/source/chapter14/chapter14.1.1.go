package main

import (
	"fmt"
	"os"
)

// 调用系统信息
func main() {
	// 获取主机名
	hn, _ := os.Hostname()
	fmt.Printf("获取主机名：%v\n", hn)
	// 获取用户ID
	fmt.Printf("获取用户ID：%v\n", os.Getuid())
	// 获取有效用户ID
	fmt.Printf("获取有效用户ID：%v\n", os.Geteuid())
	// 获取组ID
	fmt.Printf("获取组ID：%v\n", os.Getgid())
	// 获取有效组ID
	fmt.Printf("获取有效组ID：%v\n", os.Getegid())
	// 获取进程ID
	fmt.Printf("获取进程ID：%v\n", os.Getpid())
	// 获取父进程ID
	fmt.Printf("获取父进程ID：%v\n", os.Getppid())
	// 获取某个环境变量的值
	fmt.Printf("获取环境变量的值：%v\n", os.Getenv("GOPATH"))
	// 设置某个环境变量的值
	os.Setenv("TEST", "test")
	// 删除某个环境变量
	os.Unsetenv("TEST")
	// 获取所有环境变量
	for _, e := range os.Environ() {
		fmt.Printf("环境变量：%v\n", e)
	}
	// 获取某个环境变量
	fmt.Printf("获取GOPATH环境变量：%v\n", os.Getenv("GOPATH"))
	// 删除所有环境变量
	// os.Clearenv()
}
