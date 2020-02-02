package main

import "fmt"

func main() {
	var i1 = 101
	//十进制
	fmt.Printf("%d\n", i1)
	fmt.Printf("%o\n", i1) //十进制转八进制
	fmt.Printf("%b\n", i1) //十进制转二进制
	fmt.Printf("%x\n", i1) //十进制转十六进制
	//八进制
	i2 := 077
	fmt.Printf("%d\n", i2)
	//十六进制
	i3 := 0x1234567
	fmt.Printf("%d\n", i3)

	//查看变量类型

	fmt.Printf("%T\n", i3)
	//声明ini8 类型
	i4 := int(8) //明确指定int8类型 否则为int类型
	fmt.Printf("%T\n", i4)
}
