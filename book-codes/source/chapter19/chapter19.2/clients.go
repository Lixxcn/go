package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"time"
)

func onMessageReceived(conn *net.TCPConn) {
	// 创建TCP连接对象的IO
	reader := bufio.NewReader(conn)
	// 发送数据
	b := []byte(conn.LocalAddr().String() + "客户端在发送数据.\n")
	conn.Write(b)
	for {
		// 获取TCP连接对象的数据流
		msg, err := reader.ReadString('\n')
		fmt.Printf("客户端收到服务端数据：%v\n", msg)
		// 出现异常终止循环
		if err != nil || err == io.EOF {
			break
		}
		time.Sleep(time.Second * 2)
		// 通过TCP连接对象发送数据给服务端
		// 将数据转为字节类型的切片
		b := []byte(conn.LocalAddr().String() + "客户端在发送数据.\n")
		// Write()发送数据
		_, err = conn.Write(b)
		if err != nil {
			break
		}
	}
}

func main() {
	// 创建TCP结构体对象
	var tcpAddr *net.TCPAddr
	// 实例化结构体TCPAddr
	tcpAddr, _ = net.ResolveTCPAddr("tcp", "127.0.0.1:9999")
	// 创建TCP连接对象，连接TCP服务端
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		fmt.Printf("客户端连接错误：%v\n", err.Error())
		return
	}
	// 关闭连接
	defer conn.Close()
	fmt.Printf("客户端连接成功：%v\n", conn.LocalAddr().String())
	onMessageReceived(conn)
}
