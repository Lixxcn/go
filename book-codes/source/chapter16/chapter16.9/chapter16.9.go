package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

// 定义结构体
type User struct {
	gorm.Model
	ID         uint   `gorm:"primary_key"`
	Name       string `gorm:"type:varchar(255)"`
	Age        string `gorm:"type:varchar(255)"`
	Profession string `gorm:"type:varchar(255)"`
	OrganizeID int
	// 设置OrganizeID为外键
	Organize Organize `gorm:"ForeignKey:OrganizeID"`
}

type Organize struct {
	gorm.Model
	ID    uint   `gorm:"primary_key"`
	Name  string `gorm:"type:varchar(255)"`
	Site  string `gorm:"type:varchar(255)"`
	Grade string `gorm:"type:varchar(255)"`
}

func connect_db() *gorm.DB {
	// 使用gorm.io/driver/mysql连接数据库
	dsn := `root:1234@tcp(127.0.0.1:3306)/enterprise?charset=utf8mb4&parseTime=True&loc=Local`
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	sqlDB, _ := db.DB()
	// 设置连接池
	// SetMaxIdleConns设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns设置打开数据库连接的最大数量
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime设置连接可复用的最大时间
	sqlDB.SetConnMaxLifetime(time.Hour)
	// 执行数据迁移，创建数据表
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Organize{})
	return db
}

func main() {
	db := connect_db()
	// 系统功能
	for {
		var s int
		fmt.Printf("欢迎来到企业员工信息管理系统\n")
		fmt.Printf("员工管理请按1,组织架构管理请按2,退出请按3：\n")
		fmt.Scanln(&s)
		if s == 1 {
			var u int
			fmt.Printf("新增员工请按1,删除员工请按2,查询员工请按3：\n")
			fmt.Scanln(&u)
			if u == 1 {
				var name, age, profession string
				var organizeID int
				fmt.Printf("请输入名字：\n")
				fmt.Scanln(&name)
				fmt.Printf("请输入年龄：\n")
				fmt.Scanln(&age)
				fmt.Printf("请输入职位：\n")
				fmt.Scanln(&profession)
				fmt.Printf("请输入所属组织编号：\n")
				fmt.Scanln(&organizeID)
				user := User{Name: name, Age: age, Profession: profession, OrganizeID: organizeID}
				db.Create(&user)
				fmt.Printf("新员工%v添加成功：\n", name)
			} else if u == 2 {
				var name string
				fmt.Printf("请输入需要删除的名字：\n")
				fmt.Scanln(&name)
				db.Where("name = ?", name).Delete(&User{})
				fmt.Printf("员工%v删除成功：\n", name)
			} else if u == 3 {
				// 查询所有员工信息
				var ls []User
				db.Preload("Organize").Find(&ls)
				for _, v := range ls {
					fmt.Printf("员工%v的职位:%v,所属组织:%v\n\n", v.Name, v.Profession, v.Organize.Name)
				}
			}
		} else if s == 2 {
			// 课后作业
		} else if s == 3 {
			return
		} else {
			fmt.Printf("请按照提示输入")
		}
	}
}
