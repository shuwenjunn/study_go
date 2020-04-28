package main

import "fmt"

// 接口的实现

type animal interface {
	move()
	eat(string)
}

type cat struct {
	name string
	feet int8
}

func (c *cat) move() {
	fmt.Println("走猫步")
}

func (c *cat) eat(food string) {
	fmt.Println("猫吃鱼", food)
}

func main() {
	var a1 animal
	c1 := cat{
		name: "tom",
		feet: 4,
	}
	c2 := &cat{
		name: "tom2",
		feet: 4,
	}

	a1 = &c1
	a1 = c2

	fmt.Println(a1)

}
