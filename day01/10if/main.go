package main

import "fmt"

func main() {
	age := 19

	if age > 18 {
		fmt.Println("成年了")
	} else {
		fmt.Println("未成年")
	}

	if age > 6 {
		fmt.Println("一个小学生")
	} else if age > 18 {
		fmt.Println("少年")
	} else {
		fmt.Println("其他类型")
	}
}
