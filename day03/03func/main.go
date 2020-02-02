package main

import "fmt"

func f1() {
	fmt.Println("hello,沙河！")
}

func f2(name string) {
	fmt.Println(name)
}

func f3(x int, y int) int {
	return x + y
}

//  参数 类型简写
func f4(x, y int) int {
	return x + y
}

// 可变参数
func f5(x int, y ...int) int {
	fmt.Println(y)
	//y 是一个int 类型的切片
	return 1
}

// 命名返回值 在函数中直接使用命名过的遍历 return 后面的变量也可以省略
func f6(x, y int) (sum int) {
	sum = x + y
	return
}

// go 语言中多个返回值
func f7(x, y int) (sum, sub int) {
	sum = x + y
	sub = x - y
	return
}

func main() {
	f1()
	f2("ni")
	fmt.Println(f3(100, 200))

	f5(1, 23, 4, 523)
	fmt.Println(f6(1, 1))
	fmt.Println(f7(1, 1))
}
