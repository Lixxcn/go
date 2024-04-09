package main

import (
	"fmt"
	"io"
	"net"
	"time"
)

func main() {
	// 创建UDP结构体对象
	var udpAddr *net.UDPAddr
	// 实例化结构体udpAddr
	udpAddr, _ = net.ResolveUDPAddr("udp", "127.0.0.1:9999")
	// 创建UDP连接对象，连接UDP服务端
	conn, err := net.DialUDP("udp", nil, udpAddr)
	if err != nil {
		fmt.Printf("客户端连接错误：%v\n", err.Error())
		return
	}
	fmt.Printf("客户端连接成功：%v\n", conn.LocalAddr().String())
	for {
		// 通过UDP连接对象发送数据给服务端
		// 将数据转为字节类型的切片
		b := []byte(conn.LocalAddr().String() + "--客户端在发送数据")
		// Write()发送数据
		conn.Write(b)
		time.Sleep(time.Second * 2)
		// 获取UDP连接对象的数据流
		// 接收数据
		message := make([]byte, 4096)
		// ReadFromUDP()设有3个返回值
		// 第一个返回值是数据长度
		// 第二个返回值是服务端IP地址
		// 第三个返回值是异常信息
		n, _, err := conn.ReadFromUDP(message)
		// string(message[:n])根据数据长度截取数据
		fmt.Printf("客户端收到数据：%v\n", string(message[:n]))
		// 出现异常终止循环
		if err != nil || err == io.EOF {
			break
		}
	}
}
