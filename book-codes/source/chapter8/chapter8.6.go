package main

import "fmt"

func main() {
	type cars struct {
		name  string
		price int
	}

	type person struct {
		name   string
		age    int
		cars   cars
		hourse struct {
			name  string
			price int
		}
	}

	c := cars{name: "BWM", price: 500000}
	p := person{name: "Tim", age: 30, cars: c}
	fmt.Printf("个人名称：%v\n", p.name)
	fmt.Printf("个人年龄：%v\n", p.age)
	fmt.Printf("个人拥有车辆：%v\n", p.cars.name)
	fmt.Printf("车辆价钱为：%v\n", p.cars.price)
}
