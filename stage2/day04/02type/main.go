package main

import "fmt"

//自定义类型和类型别名

type myInt int

type youInt = int //类型别名

func main() {
	var n myInt
	var m youInt
	n = 100
	m = 100
	fmt.Printf("%T\n", n)
	fmt.Printf("%T\n", m)

	var c rune // rune 是int32类型别名
	c = '中'
	fmt.Printf("%T\n", c)
}
