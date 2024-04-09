package main

import "fmt"

func main() {
	//// 使用方法1
	//finger := 3
	//switch finger {
	//	// 当变量finger=1的时候
	//	case 1:
	//		fmt.Println("大拇指")
	//	// 当变量finger=2的时候
	//	case 2:
	//		fmt.Println("食指")
	//	// 当变量finger=3的时候
	//	case 3:
	//		fmt.Println("中指")
	//	case 4:
	//		fmt.Println("无名指")
	//	case 5:
	//		fmt.Println("小拇指")
	//	default:
	//		fmt.Println("无效的输入！")
	//}
	//
	//// 使用方法2
	//rand.Seed(time.Now().Unix())
	//num := rand.Intn(100)
	//switch {
	//	// 判断num是否大于20
	//	case num < 20:
	//		fmt.Printf("变量num的值为：%v，小于20\n", num)
	//	default:
	//		fmt.Printf("变量num的值为：%v，大于20\n", num)
	//}
	//
	//// 使用方法3
	//rand.Seed(time.Now().Unix())
	//switch n := rand.Intn(9); n {
	//	// 变量n在（1, 3, 5, 7, 9）区间内
	//	case 1, 3, 5, 7, 9:
	//		fmt.Printf("奇数，值为：%v，小于20\n", n)
	//	// 变量n在（2, 4, 6, 8）区间内
	//	case 2, 4, 6, 8:
	//		fmt.Printf("偶数，值为：%v，小于20\n", n)
	//	default:
	//		fmt.Printf("啥也不是")
	//}
	finger := 1
	switch finger {
	// 当变量finger=1的时候
	case 1:
		fmt.Println("大拇指")
		fallthrough
	// 当变量finger=2的时候
	case 2:
		fmt.Println("食指")
		fallthrough
	// 当变量finger=3的时候
	case 3:
		fmt.Println("中指")
	case 4:
		fmt.Println("无名指")
	case 5:
		fmt.Println("小拇指")
	default:
		fmt.Println("无效的输入！")
	}
}
