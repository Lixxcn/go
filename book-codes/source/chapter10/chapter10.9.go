package main

import (
	"fmt"
	"reflect"
)

// 定义结构体
type cat struct {
	Name string
	Age  int `json:"age" id:"101"`
}

func main() {
	// 字符串
	s := "golang"
	vs := reflect.ValueOf(&s)
	fmt.Printf("反射字符串指针的内存：%v，数值：%v\n", vs, vs.Elem())
	vs.Elem().SetString("hello")
	fmt.Printf("反射字符串指针的内存：%v，数值：%v\n", vs, vs.Elem())

	// 切片
	sli := []interface{}{1, 2, "Go"}
	vsli := reflect.ValueOf(&sli)
	fmt.Printf("反射切片指针的内存：%v，数值：%v\n", vsli, vsli.Elem())
	vsli.Elem().Index(0).Set(reflect.ValueOf("golang"))
	fmt.Printf("反射切片指针的内存：%v，数值：%v\n", vsli, vsli.Elem())

	// 集合
	m := make(map[string]interface{})
	m["name"] = "Tim"
	vm := reflect.ValueOf(&m)
	fmt.Printf("反射集合指针的内存：%v，数值：%v\n", vm, vm.Elem())
	rv := reflect.ValueOf("name")
	rk := reflect.ValueOf("golang")
	vm.Elem().SetMapIndex(rv, rk)
	fmt.Printf("反射集合指针的内存：%v，数值：%v\n", vm, vm.Elem())

	// 结构体
	c := cat{Name: "Lily", Age: 18}
	vcp := reflect.ValueOf(&c).Pointer()
	fmt.Printf("反射结构体指针的内存：%v，数值：%v\n", vcp, vcp)
	vc := reflect.ValueOf(&c)
	fmt.Printf("反射结构体指针的内存：%v，数值：%v\n", vc, vc.Elem())
	vc.Elem().FieldByName("Name").Set(reflect.ValueOf("Tom"))
	fmt.Printf("反射结构体指针的内存：%v，数值：%v\n", vc, vc.Elem())

	// 指针
	var prt *string
	name := "point"
	// 给指针赋于变量name的内存地址
	prt = &name
	vpp := reflect.ValueOf(&prt)
	fmt.Printf("反射指针的指针内存：%v，数值：%v\n", vpp, vpp.Elem())
	vp := reflect.ValueOf(prt)
	fmt.Printf("反射指针的内存：%v，数值：%v\n", vp, vp.Elem())
	// 使用反射创建新指针
	nprt := reflect.New(reflect.TypeOf(*prt))
	fmt.Printf("新反射指针的内存：%v，数值：%v\n", nprt, nprt.Elem())
	nprt.Elem().Set(reflect.ValueOf(name))
	fmt.Printf("新反射指针的内存：%v，数值：%v\n", nprt, nprt.Elem())
	// 使用反射创建新指针的指针
	nprtt := reflect.New(reflect.TypeOf(prt))
	fmt.Printf("新反射指针的指针内存：%v，数值：%v\n", nprtt, nprtt.Elem())
	nprtt.Elem().Set(reflect.ValueOf(&name))
	fmt.Printf("新反射指针的指针内存：%v，数值：%v\n", nprtt, nprtt.Elem())
}
