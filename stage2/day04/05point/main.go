package main

import "fmt"

func main() {
	var a int
	a = 100
	b := &a

	fmt.Printf("type  a:%T,type  b:%T\n", a, b)

	//将a的内存地址拿出来
	fmt.Printf("%p\n", &a)
	fmt.Printf("%p\n", b)
	fmt.Printf("%p\n", &b) //取出变量b的 内存地址

}
