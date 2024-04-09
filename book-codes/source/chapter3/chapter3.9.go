package main

import (
	"fmt"
)

func main() {
	begin := `
	欢迎致电中国移动10086客服服务热线，业务查询请按1，手机充值请按2，
	业务办理请按3，语音导航请按4，人工服务请按0`
	fmt.Println(begin)
	// 定义变量
	var service int
	// 输出服务提示
	fmt.Println("请选择您的服务内容：")
	fmt.Scanln(&service)
	// 判断服务选择
	if service == 1 {
		// 业务查询
		// 输出业务查询的服务提示
		fmt.Println("话费查询请按1，流量查询请按2，套餐业务查询请按3")
		var num int
		fmt.Scanln(&num)
		if num == 1 {
			fmt.Println("您的话费余额为100元。")
		} else if num == 2 {
			fmt.Println("您的流量剩余100MB。")
		} else if num == 3 {
			fmt.Println("您的当前套餐为XXX。")
		}
	} else if service == 2 {
		// 手机充值
		// 输出手机充值的服务提示
		var codeNum string
		code := "123abc"
		fmt.Println("请输入充值卡的密码并按#键结束：")
		fmt.Scanln(&codeNum)
		// 将用户输入数据截取，去掉#
		if code == codeNum[:len(codeNum)-1] {
			fmt.Println("充值成功，您的余额为120元")
		} else {
			fmt.Println("充值失败，请输入正确的充值密码。")
		}
	} else if service == 3 {
		// 业务办理
	} else if service == 4 {
		// 语音导航
	} else if service == 0 {
		// 人工服务
	}
}
