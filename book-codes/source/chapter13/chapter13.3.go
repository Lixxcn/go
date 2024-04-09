package main

import (
	// 导入math/rand包
	"math/rand"
	// 导入crypto/rand，将包改名为crand
	crand "crypto/rand"
)

func main() {
	// 调用math/rand包的函数Int()
	rand.Int()
	// 调用crypto/rand包的函数Read()
	crand.Read([]byte{'a', 'b'})
}
