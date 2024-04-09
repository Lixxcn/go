package main

import (
	"html/template"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	// 定义变量，用于传递给模版文件
	value := "This is route"
	// 创建tmpl.html的模版对象
	t, _ := template.ParseFiles("t1.html", "t2.html")
	// 执行模版解析
	t.Execute(w, value)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/", index)
	server.ListenAndServe()
}
