package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// 如果当前路径没有MyDb.db，程序会自动创建
	db, _ := sql.Open("sqlite3", "MyDb.db")

	// 通过程序执行SQL语句创建数据表
	sql_table := `CREATE TABLE IF NOT EXISTS "userinfo" (
	   "id" INTEGER PRIMARY KEY AUTOINCREMENT,
	   "username" VARCHAR(64) NULL,
	   "age" INT(10) NULL,
	   "created" TIMESTAMP default (datetime('now','localtime'))
				)`
	// 执行SQL语句
	db.Exec(sql_table)

	// 新增数据
	stmt, _ := db.Prepare("INSERT INTO userinfo(username, age) values(?, ?)")
	// 传递参数并执行SQL语句
	res, _ := stmt.Exec("Tom", "18")
	// 返回新增数据的ID
	id, _ := res.LastInsertId()
	fmt.Printf("新增数据的ID：%v\n", id)

	// 批量新增数据
	UserList := [][]interface{}{{"Lily", 22}, {"Jim", 30}}
	for _, i := range UserList {
		// 新增数据
		stmt, _ := db.Prepare("INSERT INTO userinfo(username, age) values(?, ?)")
		// 传递参数并执行SQL语句
		res, _ := stmt.Exec(i[0], i[1])
		// 返回新增数据的ID
		id, _ := res.LastInsertId()
		fmt.Printf("批量新增数据的ID：%v\n", id)
	}

	// 更新数据
	stmt, _ = db.Prepare("update userinfo set username=? where id=?")
	// 传递参数并执行SQL语句
	res, _ = stmt.Exec("Tim", 1)
	// 受影响的数据行数，返回int64类型数据
	affect, _ := res.RowsAffected()
	fmt.Printf("更新数据受影响的数据行数：%v\n", affect)

	// 批量更新数据
	UserList1 := [][]interface{}{{"Betty", 3}, {"Jon", 4}}
	for _, i := range UserList1 {
		stmt, _ := db.Prepare("update userinfo set username=? where id=?")
		// 传递参数并执行SQL语句
		res, _ := stmt.Exec(i[0], i[1])
		// 受影响的数据行数，返回int64类型数据
		affect, _ := res.RowsAffected()
		fmt.Printf("更新数据受影响的数据行数：%v\n", affect)
	}

	// 删除数据
	stmt, _ = db.Prepare("delete from userinfo where id=?")
	// 将想删除的id输入进去就可以删除输入的id
	res, _ = stmt.Exec(1)
	// 受影响的数据行数，返回int64类型数据
	affect, _ = res.RowsAffected()
	fmt.Printf("删除数据受影响的数据行数：%v\n", affect)

	// 批量删除数据
	IDList := []int{3, 4}
	for _, i := range IDList {
		// 通过循环删除多条数据，每次循环删除一条数据
		stmt, _ := db.Prepare("delete from userinfo where id=?")
		res, _ := stmt.Exec(i)
		// 受影响的数据行数，返回int64类型数据
		affect, _ := res.RowsAffected()
		fmt.Printf("批量删除数据受影响的数据行数：%v\n", affect)
	}

	// 查询数据
	rows, _ := db.Query("SELECT * FROM userinfo where id=?", 2)
	// 遍历所有查询结果
	var ids, age int
	var un, ct string
	for rows.Next() {
		rows.Scan(&ids, &un, &age, &ct)
		fmt.Printf("当前数据：%v,%v,%v,%v\n", ids, un, age, ct)
	}
	// 关闭数据库连接
	db.Close()
}
