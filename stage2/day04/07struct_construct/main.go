package main

import "fmt"

type person struct {
	name string
	age  int
}

//构造函数 约定成俗 以new 开头的函数

func newPerson(name string, age int) *person {
	return &person{
		name: name,
		age:  age,
	}
}

func main() {
	p1 := newPerson("袁栓", 18)
	p2 := newPerson("zhou", 9000)

	fmt.Println(p1.age)
	fmt.Println(*p2)
}
