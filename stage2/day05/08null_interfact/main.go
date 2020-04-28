package main

import "fmt"

func show(a interface{}) {
	fmt.Println(a)
}

func main() {
	// 空接口类型
	var m1 map[string]interface{}
	m1 = make(map[string]interface{}, 16)
	m1["name"] = "zhoulin"
	m1["age"] = 19
	m1["gender"] = true
	fmt.Println(m1)

	show(18)
	show("你妈")
	show(m1)
}
