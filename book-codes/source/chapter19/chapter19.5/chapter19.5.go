package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func body(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// 参数w是响应对象
	// 参数r是请求对象
	// 参数p是路由变量
	w.Write([]byte("This is body"))
}

func user(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// 参数w是响应对象
	// 参数r是请求对象
	// 参数p是路由变量
	v := "This is user, name is " + p.ByName("name")
	w.Write([]byte(v))
}

func main() {
	// 定义路由与路由的处理函数body
	router := httprouter.New()
	// 设置HTTP的GET请求
	router.GET("/", body)
	// 设置HTTP的POST请求
	router.POST("/", body)
	// 设置HTTP的PUT请求
	router.PUT("/", body)
	// 设置HTTP的DELETE请求
	router.DELETE("/", body)
	// 匹配内容直到下一个斜线/或者路径的结尾
	router.GET("/user1/:name", user)
	// 从指定位置开始匹配到结尾
	// 不能在根目录下直接使用，如/*name则提示异常
	router.GET("/user2/*name", user)
	// 将httprouter绑定在net/http
	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: router,
	}
	// 运行服务
	server.ListenAndServe()

}
