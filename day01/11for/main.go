package main

import "fmt"

func main() {
	//基本格式
	//for i := 0; i < 10; i++ {
	//	fmt.Println(i)
	//}

	//变种i
	//i := 5
	//for ; i < 10; i++ {
	//	fmt.Println(i)
	//}

	//变种ii
	//i := 5
	//for ; i < 10; {
	//	fmt.Println(i)
	//	i++
	//}

	//无限循环
	//for{
	//	fmt.Println("无限循环")
	//}

	//for range 循环
	s := "hello沙河"
	for i, v := range s {
		fmt.Printf("%d %c\n", i, v)
	}
}
