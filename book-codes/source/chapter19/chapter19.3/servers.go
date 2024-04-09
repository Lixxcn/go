package main

import (
	"fmt"
	"io"
	"net"
	"time"
)

func main() {
	// 定义UDP对象
	var udpAddr *net.UDPAddr
	// UDP对象绑定IP和端口
	udpAddr, _ = net.ResolveUDPAddr("udp", "127.0.0.1:9999")
	// 创建UDP连接对象
	conn, _ := net.ListenUDP("udp", udpAddr)
	defer conn.Close()
	// 接收并返回消息
	for {
		// 获取接收数据
		message := make([]byte, 4096)
		// ReadFromUDP()设有3个返回值
		// 第一个返回值是数据长度
		// 第二个返回值是客户端IP地址
		// 第三个返回值是异常信息
		n, addr, err := conn.ReadFromUDP(message)
		// 出现异常说明连接异常
		if err != nil || err == io.EOF {
			break
		}
		// string(message[:n])根据数据长度截取数据
		fmt.Printf("服务端接收数据：%v\n", string(message[:n]))
		time.Sleep(time.Second * 3)
		// 发送数据
		msg := conn.LocalAddr().String() + "--服务端发送数据"
		b := []byte(msg)
		// 服务器发送数据必须调用WriteToUDP()
		// addr是客户端的IP地址
		// WriteToUDP()根据指定IP地址发送数据
		conn.WriteToUDP(b, addr)
	}
}
