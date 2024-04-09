package main

import "fmt"

func main() {
	// 在最外层循环中设置标签，标签名自行修改
for1:
	for i := 0; i < 3; i++ {
		for k := 1; k < 10; k++ {
			fmt.Printf("%v:%v\n", i, k)
			// break后加上标签名，直接终止最外层循环
			break for1
		}
	}
}
