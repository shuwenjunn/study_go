package main

import "fmt"

func f1() {
	fmt.Println("hello 沙河")
}

func f2() int {
	return 10
}

//函数也可以作为参数的类型
func f3(x func() int) {
	ret := x()
	print(ret)
}

func f4(x, y int) int {
	return x + y
}

// 函数还可以作为返回值

func f5(x func()) func() {
	return x
}

func main() {

	a := f1
	fmt.Printf("%T\n", a)

	b := f2
	fmt.Printf("%T\n", b)

	f3(f2)
	fmt.Printf("%T\n", f4)

}
