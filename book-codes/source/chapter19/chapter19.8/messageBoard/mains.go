package main

import (
	"github.com/julienschmidt/httprouter"
	"html/template"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// panic用于测试500页面
	// panic("这是自定义异常")
	// 判断当前请求方式
	if r.Method == "POST" {
		// 从表单获取请求参数name和content
		name := r.PostFormValue("name")
		content := r.PostFormValue("content")
		// 将请求参数写入数据表
		Db.Create(&Message{Name: name, Content: content})
		// 重定向当前路由，以GET请求方式返回给用户
		w.Header().Set("Location", "/")
		w.WriteHeader(http.StatusMovedPermanently)
	}
	// 查询数据表所有数据
	var data []Message
	Db.Model(&Message{}).Scan(&data)
	// 创建html的模版对象
	t, _ := template.ParseFiles("templates/base.html", "templates/index.html")
	// 执行模版解析
	t.Execute(w, data)
}

func MyNotFound(w http.ResponseWriter, r *http.Request) {
	// 创建html的模版对象
	t, _ := template.ParseFiles("templates/base.html", "templates/404.html")
	// 执行模版解析
	t.Execute(w, "")
}

func MyPanic(w http.ResponseWriter, r *http.Request, e interface{}) {
	// 参数e代表程序运行的异常信息
	// 设置HTTP状态码500
	w.WriteHeader(http.StatusInternalServerError)
	// 创建html的模版对象
	t, _ := template.ParseFiles("templates/base.html", "templates/500.html")
	// 执行模版解析
	t.Execute(w, "")
}

func main() {
	router := httprouter.New()
	// 定义路由
	router.GET("/", index)
	router.POST("/", index)
	// 配置静态文件路径
	router.ServeFiles("/static/*filepath", http.Dir("./static"))
	// 自定义404
	router.NotFound = http.HandlerFunc(MyNotFound)
	// 自定义所有异常处理
	router.PanicHandler = MyPanic
	// 将httprouter绑定在net/http
	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: router,
	}
	// 运行服务
	server.ListenAndServe()
}
