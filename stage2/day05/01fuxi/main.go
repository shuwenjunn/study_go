package main

import "fmt"

type animal struct {
	name  string
	cate  map[string]string // 动物的类别里面有多少
	count int
}

//给猫的数量加一

func (a *animal) addCount() {
	a.count++
}

func (a animal) addCate() { //引用类型的不需要取内存地址
	a.cate["短猫"] = "英国"
}

func main() {
	// 结构体复习 何时需要指针 何时不需要

	cat := animal{
		name:  "猫咪",
		cate:  map[string]string{"加菲猫": "美国", "黑猫": "中国"},
		count: 5,
	}

	fmt.Println(cat)
	cat.addCate()
	fmt.Println(cat)
}
