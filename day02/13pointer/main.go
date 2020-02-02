package main

import "fmt"

//指针

func main() {
	//去地址
	n := 18
	fmt.Println(&n)
	p := &n
	fmt.Println(p)
	q := *p
	fmt.Println(q)
	//根据地址取值
}
