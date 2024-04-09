package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"time"
)

func tcpPipe(conn *net.TCPConn) {
	// TCP连接的地址
	ipStr := conn.RemoteAddr().String()
	// 关闭连接
	defer func() {
		fmt.Printf("%v失去连接\n: ", ipStr)
		conn.Close()
	}()
	// 获取TCP连接对象的数据流
	reader := bufio.NewReader(conn)
	// 接收并返回消息
	for {
		// 获取接收数据
		message, err := reader.ReadString('\n')
		// 出现异常说明连接异常
		if err != nil || err == io.EOF {
			break
		}
		fmt.Printf("服务端接收数据%v\n", message)
		time.Sleep(time.Second * 3)
		// 发送数据
		msg := conn.RemoteAddr().String() + "--服务端发送数据\n"
		b := []byte(msg)
		conn.Write(b)
	}
}

func main() {
	// 定义TCP对象
	var tcpAddr *net.TCPAddr
	// TCP对象绑定IP和端口
	tcpAddr, _ = net.ResolveTCPAddr("tcp", "127.0.0.1:9999")
	// 创建TCP监听对象
	tcpListener, _ := net.ListenTCP("tcp", tcpAddr)
	// 关闭TCP监听对象
	defer tcpListener.Close()
	// 循环接收客户端的连接，创建协程去处理连接
	for {
		// 通过TCP监听对象获取与客户端的TCP连接对象
		tcpConn, err := tcpListener.AcceptTCP()
		if err != nil {
			fmt.Println(err)
			continue
		}
		// 连接成功后创建协程去处理连接
		go tcpPipe(tcpConn)
	}
}
