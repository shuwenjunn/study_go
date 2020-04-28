package main

import "fmt"

// 类型断言
func assign(a interface{}) {
	fmt.Println("%T\n", a)
	str, ok := a.(string) //猜测是string 类型 ，然后将它转成string
	if !ok {
		fmt.Println("猜错了")
	} else {
		fmt.Println(str)
	}

}

func main() {
	assign(18)
}
