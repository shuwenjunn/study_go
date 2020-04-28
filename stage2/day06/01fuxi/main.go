package main

import "fmt"

// 类型断言

func main() {
	var a interface{}

	a = 100

	_, ok := a.(int)

	if ok {
		fmt.Println("猜对了")
	} else {
		fmt.Println("猜错了")
	}

	switch v2 := a.(type) {
	case int8:
		fmt.Println("int8", v2)
	case int16:
		fmt.Println("int16", v2)
	case int:
		fmt.Println("int", v2)
	case string:
		fmt.Println("string", v2)
	default:
		fmt.Println("啥都不是")
	}
}
