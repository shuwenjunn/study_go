package main

import "fmt"

//接口是一种 特殊的类型 他规定了变量有 哪些方法

//
type speaker interface {
	speak() //只要实现speak 方法的变量 都是speaker 类型

}

type cat struct {
}

type person struct {
}

func (c cat) speak() {
	fmt.Println("miao~")
}

func (p person) speak() {
	fmt.Println("cnm~")
}

func da(x speaker) {
	x.speak()
}

func main() {
	var c cat
	var p person

	da(c)
	da(p)
}
