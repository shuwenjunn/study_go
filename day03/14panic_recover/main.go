package main

import "fmt"

func funcA() {
	fmt.Println("a")
}

func funcB() {
	//panic("出现了严重的错误") //意外执行退出
	fmt.Println("b")
}
func funcC() {

	fmt.Println("c")
}

func main() {
	funcA()
	funcB()
	funcC()

}
