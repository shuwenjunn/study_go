package main

import "fmt"

// 函数
// 函数的定义

func sum(x int, y int) (ret int) {
	return x + y
}

// 无返回值的函数
func f1(x int, y int) {
	fmt.Println(x + y)
}

// 没有参数有返回值

func f2() int {
	return 111
}

// 参数可以命名 也可以不命名

// 多个返回值
func f5() (int, string) {
	return 1, "hhah"
}

//参数的类型简写
// 当参数中的连续类型一样时
func f6(x, y int, m, n string, a, b bool) int {
	return 123
}

// 可变长参数 y 可以传递多个值 必须放到函数参数的最后

func f7(x string, y ...int) {
	fmt.Println(x)
	fmt.Println(y)
}

// go 中没有默认参数的概念

func main() {
	a := sum(1, 2)
	fmt.Println(a)
	f1(1, 2)

	fmt.Println(f2())

	fmt.Println(f5())
	m, n := f5()
	fmt.Println(m, n)
	f7("aaa")
	f7("aaa", 1, 234, 5)
}
