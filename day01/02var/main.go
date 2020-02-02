package main

import "fmt"

//申明变量 go 推荐使用驼峰方式风命名方式 同 react

//单独声明
//var s1 string
//var age int
//var isOk bool

//批量声明

var (
	name string //""
	age  int    //0
	isOk bool   // false
)

func main() {
	name = "理想"
	age = 18
	isOk = true
	//go 语言变量使用必须使用 不使用无法编辑

	fmt.Printf("name:%s", name)
	fmt.Println(age)
	fmt.Print(isOk)

	//	申明变量 同时赋值
	var s1 string = "test"
	fmt.Println(s1)
	var s2 = "我是s2" //类型推导
	fmt.Println(s2)

	//	简短变量声明 只能在函数里面使用
	s3 := "hahhah"
	fmt.Println(s3)
}
