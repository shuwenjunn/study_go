package main

import "fmt"

// 变量的作用域  全局作用域 语句块作用域 函数作用域

var x = 100

func f1() {
	// 函数中查找变量的顺序
	//1 先在函数内部查找
	//2 找不到就向函数的外部查找，一直找到全局变量

	name := "理想" //函数定义内部变量只能在函数内部使用
	fmt.Println(x, name)
}

func main() {
	f1()
}
