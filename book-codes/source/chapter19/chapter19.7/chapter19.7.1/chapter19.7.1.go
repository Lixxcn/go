package main

import (
	"html/template"
	"net/http"
	"time"
)

func formatDate(t time.Time) string {
	return t.Format("2006-01-02")
}

func index(w http.ResponseWriter, r *http.Request) {
	// 将函数formatDate绑定template的FuncMap，命名为fdate
	funcMap := template.FuncMap{"fdate": formatDate}
	// 创建新模版tmpl.html，并将funcMap注册到模版tmpl.html
	t := template.New("tmpl.html").Funcs(funcMap)
	// 创建tmpl.html的模版对象
	t, _ = t.ParseFiles("tmpl.html")
	// 执行模版解析
	t.Execute(w, time.Now())
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/", index)
	server.ListenAndServe()
}
