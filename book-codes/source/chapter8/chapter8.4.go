//package main
//
//import (
//	"encoding/json"
//	"fmt"
//)
//
//type Student struct {
//	Name  string
//	Age   int
//	Score int
//}
//
//func main() {
//	var stu Student = Student{
//		Name:  "张三",
//		Age:   22,
//		Score: 88,
//	}
//
//	data, _ := json.Marshal(stu)
//	fmt.Printf("结构体转换JSON：%v\n", string(data))
//}

package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	// `json:"name"`表示json序列化时，结构体成员展示形式为name
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Score int    `json:"score"`
}

func main() {
	var stu Student = Student{
		Name:  "张三",
		Age:   22,
		Score: 88,
	}

	data, _ := json.Marshal(stu)
	fmt.Println(string(data))
}
