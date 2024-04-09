package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Write()接受一个字节切片作为参数
// 字节切片作为HTTP响应内容，支持HTML转义
func indexExample(w http.ResponseWriter, r *http.Request) {
	str := `<html>
			<head><title>My Go</title></head>
			<body><h1>Hello World</h1></body>
			</html>`
	w.Write([]byte(str))
}

// WriteHeader设置响应状态码
func errorExample(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(501)
	str := `<html>
			<head><title>My Go</title></head>
			<body><h1>Hello World</h1></body>
			</html>`
	w.Write([]byte(str))
}

// 在Header中设置参数Location
// 并使用WriteHeader设置302状态码即可实现URL重定向
// 重定向的URL为参数Location的参数值
func redirectExample(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Location", "https://www.baidu.com/")
	w.WriteHeader(302)
}

// 定义结构体Post，用于生成JSON数据
type Post struct {
	User    string
	Threads []string
}

// 在Header中设置参数Content-Type
// 参数值为application/json，将响应内容以JSON表示
// 使用结构体Post生成JSON数据
// 由Write方法将JSON数据作为响应内容输出
func jsonExample(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	post := &Post{
		User:    "Go",
		Threads: []string{"first", "second", "third"},
	}
	json, _ := json.Marshal(post)
	w.Write(json)
}

func cookieExample(w http.ResponseWriter, r *http.Request) {
	// 获取HTTP请求的Cookie
	c, _ := r.Cookie("csrftoken")
	// 获取Cookie某个属性值
	fmt.Printf("获取HTTP请求的Cookie:%v\n", c)
	// 设置响应内容的Cookie
	cookie := &http.Cookie{
		Name:   "sessionid",
		Value:  "lkjsdfklsjfklsfdsfdjslf",
		MaxAge: 3600,
		Domain: "localhost",
		Path:   "/",
	}
	http.SetCookie(w, cookie)
	w.Write([]byte("This is Cookie"))
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/", indexExample)
	http.HandleFunc("/error", errorExample)
	http.HandleFunc("/redirect", redirectExample)
	http.HandleFunc("/json", jsonExample)
	http.HandleFunc("/cookie", cookieExample)
	server.ListenAndServe()
}
