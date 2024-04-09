package main

import (
	// 导入内置包math/rand，并使用无包名调用
	. "math/rand"
	// 导入内置包crypto/rand
	"crypto/rand"
)

func main() {
	// 调用math/rand包的函数Int()
	Int()
	// 调用crypto/rand包的函数Read()
	rand.Read([]byte{'a', 'b'})
}
