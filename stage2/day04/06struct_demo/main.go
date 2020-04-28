package main

import "fmt"

// 结构体占用一块 连续的内存
type x struct {
	a, b, c int8 // 8bit=1byte
}

func main() {
	m := x{
		a: int8(10),
		b: int8(20),
		c: int8(30),
	}

	fmt.Print("%p\n", &m.a)
	fmt.Print("%p\n", &m.b)
	fmt.Print("%p\n", &m.c)
}
