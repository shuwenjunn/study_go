package main

import "fmt"

func main() {
	//fmt.Print("沙河")
	//fmt.Print("娜扎")
	//
	//var s string
	//
	////获取用户的输入值
	//fmt.Scan(&s)
	//fmt.Println(s)

	var (
		name  string
		age   int
		class string
	)
	fmt.Scanf("%s %d %s", &name, &age, &class)
	fmt.Println(name, age, class)
}
