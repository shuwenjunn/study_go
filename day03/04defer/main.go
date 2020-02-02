package main

import "fmt"

func deferDemo() {
	fmt.Println("start")
	defer fmt.Println("嘿嘿嘿 ") // defer 把他后面的语句延迟执行了
	defer fmt.Println("呵呵呵 ") // 后进先出  一个函数可以有多个defer语句
	defer fmt.Println("哈哈哈 ") // 后进先出 多用于函数结束之前释放资源
	fmt.Println("end")
}

func main() {
	deferDemo()
}
