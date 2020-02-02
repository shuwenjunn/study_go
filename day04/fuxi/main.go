package main

import (
	"fmt"
)

func yuanshuai(name string) {
	fmt.Println("hello", name)
}

func lixiang(f func(string), name string) {
	f(name)
}

func zhoulin() func(int, int) int {
	f := func(x int, y int) int {
		return x + y
	}
	return f
}

func main() {
	lixiang(yuanshuai, "你爸爸")

	ret := zhoulin()
	fmt.Println(ret(1, 2))
}
