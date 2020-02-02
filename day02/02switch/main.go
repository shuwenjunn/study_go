package main

import "fmt"

func main() {
	//var n = 5
	//switch n {
	//case 1:
	//	fmt.Println("我是111")
	//
	//case 2:
	//	fmt.Println("我是222")
	//
	//default:
	//	fmt.Println("我是默认值")
	//
	//}
	//变种1
	//switch n := 5; n {
	//case 1:
	//	fmt.Println("我是111")
	//
	//case 2:
	//	fmt.Println("我是222")
	//
	//default:
	//	fmt.Println("我是默认值")
	//
	//}

	//变种2
	switch n := 7; n {
	case 1, 3, 5, 7:
		fmt.Println("奇数")
	case 2, 4, 6, 8:
		fmt.Sprintln("偶数")
	}
}
