package main

import "fmt"

func f1() {
	defer func() { // 函数即将结束时执行
		recover()//收集当前的错误
		fmt.Println("放手去爱")
	}()

	panic("错误")
	fmt.Println("f1")
}

func f2() {
	fmt.Println("f2")
}

func main() {
	f1()
	f2()
}
