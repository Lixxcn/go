package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func body(w http.ResponseWriter, r *http.Request) {
	// 获取请求头信息：r.Header
	// 获取请求头的某条信息
	h := r.Header.Get("Accept-Encoding")
	fmt.Printf("请求头Accept-Encoding：%v\n", h)
	// 判断请求方式
	if r.Method == "GET" {
		// GET请求参数以application/x-www-form-urlencoded编码
		// 方法1：获取GET请求的请求参数
		r.ParseForm()
		// r.Form.Get("name")等于r.Form["name"]
		fmt.Printf("获取参数的方法1：%v\n", r.Form.Get("name"))
		fmt.Printf("获取参数的方法2：%v\n", r.URL.Query())
		fmt.Printf("获取参数的方法3：%v\n", r.FormValue("name"))
		// 返回响应内容
		fmt.Fprintln(w, "This is GET")
	} else {
		// 获取POST的请求参数
		// 分别使用Form和PostForm方法获取POST的请求参数
		// 使用Form和PostForm之前必须调用ParseForm方法
		// 接收application/x-www-form-urlencoded编码的数据
		r.ParseForm()
		//fmt.Printf("Form()获取参数：%v\n", r.Form.Get("name"))
		fmt.Printf("PostForm()获取参数：%v\n", r.PostForm.Get("name"))
		// 与GET请求的r.FormValue()是同一函数方法
		fmt.Printf("FormValue()获取参数：%v\n", r.FormValue("name"))
		// PostFormValue将PostForm的功能简化
		pfv := r.PostFormValue("name")
		fmt.Printf("PostFormValue()获取参数：%v\n", pfv)

		// 获取POST的文件数据
		// MultipartForm用于文件上传
		// 使用前调用ParseMultipartForm
		// 接收multipart/form-data编码
		// 注意：FormFile是MultipartForm的简化功能
		r.ParseMultipartForm(1024)
		fmt.Printf("MultipartForm获取文件数据：%v\n", r.MultipartForm)
		// FormFile()获取上存文件
		// 返回值file代表文件内容
		// 返回值handler代表文件信息
		file, handler, _ := r.FormFile("file")
		fmt.Printf("FormFile()获取文件数据：%v、%v\n", file, handler)

		// 接收POST的JSON数据
		// 因为JSON数据使用application/json编码
		con, _ := ioutil.ReadAll(r.Body)
		fmt.Printf("接收JSON数据：%v\n", string(con))

		// 返回响应内容
		fmt.Fprintln(w, "This is POST")
	}
}

func main() {
	// 定义路由与路由的处理函数body
	http.HandleFunc("/", body)
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	// 运行服务
	server.ListenAndServe()
}
