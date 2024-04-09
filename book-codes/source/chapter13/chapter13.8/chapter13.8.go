package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	var dataSourceName = "root:1234@(127.0.0.1:3306)/test"
	// 打开数据库
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		fmt.Println(err)
	}
	// 关闭数据库
	defer db.Close()
	// 连接数据库
	fmt.Println(db.Ping())
}
