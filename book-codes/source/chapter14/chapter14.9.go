package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// 定义结构体
type Student struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Grade string `json:"grade"`
	Major string `json:"major"`
}

func data_process(style string, s ...Student) []Student {
	// 定义结构体类型的切片
	var person []Student
	// 读取JSON文件
	if style == "read" {
		f1, _ := os.OpenFile("data.json", os.O_RDWR|os.O_CREATE, 0755)
		// 实例化结构体Decoder，实现数据读取
		data := json.NewDecoder(f1)
		// 将已读取的数据加载到切片person
		data.Decode(&person)
		f1.Close()
	}
	// 写入JSON文件
	if style == "write" {
		f2, _ := os.OpenFile("data.json", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
		// 实例化结构体Encoder，实现数据写入
		encoder := json.NewEncoder(f2)
		// 数据写入必须使用文件内容覆盖，即设置os.O_TRUNC模式，否则导致内容错乱
		encoder.Encode(s)
		f2.Close()
	}
	return person
}

func main() {
	var s int
	for {
		// 输出操作提示
		fmt.Printf("欢迎来到学生信息管理系统\n")
		fmt.Printf("查询请按1,新增请按2,删除请按3,退出请按4：\n")
		// 存储用户输入数据
		fmt.Scanln(&s)
		if s == 1 {
			// 读取JSON文件获取学生信息
			data := data_process("read")
			if len(data) == 0 {
				// JSON文件读取失败
				continue
			}
			var qs int
			fmt.Printf("查询全部请按1,查询某个学生请按2：\n")
			fmt.Scanln(&qs)
			if qs == 1 {
				// 查询全部学生信息
				for _, v := range data {
					fmt.Printf("学号：%v、", v.Id)
					fmt.Printf("姓名：%v、", v.Name)
					fmt.Printf("年龄：%v、", v.Age)
					fmt.Printf("年级：%v、", v.Grade)
					fmt.Printf("专业：%v\n", v.Major)
				}
			} else if qs == 2 {
				// 查询某个学生信息
				var id int
				fmt.Printf("请输入学号查询\n")
				fmt.Scanln(&id)
				for _, v := range data {
					if v.Id == id {
						fmt.Printf("学号：%v、", v.Id)
						fmt.Printf("姓名：%v、", v.Name)
						fmt.Printf("年龄：%v、", v.Age)
						fmt.Printf("年级：%v、", v.Grade)
						fmt.Printf("专业：%v\n", v.Major)
					}
				}
			}
		} else if s == 2 {
			var id, age int
			var name, grade, major string
			fmt.Printf("请输入学号\n")
			fmt.Scanln(&id)
			fmt.Printf("请输入姓名\n")
			fmt.Scanln(&name)
			fmt.Printf("请输入年龄\n")
			fmt.Scanln(&age)
			fmt.Printf("请输入年级\n")
			fmt.Scanln(&grade)
			fmt.Printf("请输入专业\n")
			fmt.Scanln(&major)
			// 读取JSON文件获取学生信息
			data := data_process("read")
			// 实例化结构体Student
			stu := Student{
				Id: id, Name: name, Age: age,
				Grade: grade, Major: major,
			}
			// 将实例化对象写入切片data
			data = append(data, stu)
			// 将切片data写入JSON文件，利用切片的解包原理
			data_process("write", data...)
		} else if s == 3 {
			var id int
			var new_data []Student
			// 读取JSON文件获取学生信息
			data := data_process("read")
			fmt.Printf("输入学号删除学生信息：\n")
			fmt.Scanln(&id)
			fmt.Printf("删除前的学生信息：%v\n", data)
			for _, v := range data {
				if v.Id != id {
					new_data = append(new_data, v)
				}
			}
			data_process("write", new_data...)
			fmt.Printf("删除后的学生信息：%v\n", new_data)
		} else if s == 4 {
			break
		}
	}
}
