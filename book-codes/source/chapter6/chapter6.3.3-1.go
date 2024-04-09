package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// 定义字符串，用于记录JSON数据
	var j string
	j = `{"infos": [{"name": "Tom", "age": 15}, {"name": "Lily", "age": 20}]}`
	// 定义集合，value的数据类型为接口interface类型
	var m1 = map[string]interface{}{}
	// 将JSON字符串转换为集合
	json.Unmarshal([]byte(j), &m1)
	// 遍历输出JSON
	for k, v := range m1 {
		fmt.Printf("集合m1的键为：%v\n", k)
		fmt.Printf("集合m1的值为：%v\n", v)
		// 解析JSON里面的数组
		vv := v.([]interface{})
		for i := 0; i < len(vv); i++ {
			fmt.Printf("数组vv的值为：%v\n", vv[i])
			// 解析数组里面的集合
			vvv := vv[i].(map[string]interface{})
			name := vvv["name"]
			age := vvv["age"]
			fmt.Printf("键为name的数据为：%v\n", name)
			fmt.Printf("键为age的数据为：%v\n", age)
		}
	}
}
