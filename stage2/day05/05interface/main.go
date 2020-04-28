package main

import "fmt"

// 接口的实现

type animal interface {
	move()
	eat(string)
	//hah()
}

type cat struct {
	name string
	feet int8
}

type chicken struct {
	feet int8
}

func (c cat) move() {
	fmt.Println("走猫步")
}

func (c cat) eat(food string) {
	fmt.Println("猫吃鱼", food)
}

func (c chicken) move() {
	fmt.Println("激动")
}

func (c chicken) eat(food string) {
	fmt.Println("jiji", food)
}

func main() {
	var a1 animal //定义一个接口类型的变量

	bc := cat{ //定义一个cat 类型的变量
		name: "蓝猫",
		feet: 4,
	}

	a1 = bc
	fmt.Println(a1)
	a1.eat("dafen")
	fmt.Printf("%T,%T\n", a1, bc)

	kfc := chicken{feet: 2}
	a1 = kfc

	a1.eat("fad")
}
