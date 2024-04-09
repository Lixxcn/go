package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	n := "hello world"
	m := "I am Tom"
	// 使用"+"拼接字符串
	j := n + "," + m
	fmt.Println(j)
	// 使用fmt.Sprintf()拼接字符串
	k := fmt.Sprintf("%s,%s", n, m)
	fmt.Println(k)
	// 使用strings.Join()拼接字符串
	g := strings.Join([]string{n, m}, ",")
	fmt.Println(g)
	// 使用builder.WriteString()连接字符串
	var builder strings.Builder
	builder.WriteString(n)
	builder.WriteString(",")
	builder.WriteString(m)
	fmt.Println(builder.String())
	// 使用buffer.WriteString()连接字符串
	var buffer bytes.Buffer
	buffer.WriteString(n)
	buffer.WriteString(",")
	buffer.WriteString(m)
	fmt.Println(buffer.String())
}
