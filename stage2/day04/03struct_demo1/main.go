package main

import "fmt"

// 结构体
type person struct {
	name   string
	age    int
	gender string
	hobby  []string
}

func main() {
	var p person
	p.name = "舒文俊"
	p.age = 18
	p.gender = "男"
	p.hobby = []string{"足球", "篮球", "双色球"}

	fmt.Println(p, p.name, p.gender, p.age, p.hobby)
	fmt.Printf("%T\n", p)

	// 匿名结构体  多用于临时场景
	var s struct {
		x string
		y int
	}
	s.x = "嘿嘿嘿"
	s.y = 100

	fmt.Printf("%T\n", s)
}
