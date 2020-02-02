package main

import "fmt"

func main() {
	s := "hello沙河"
	n := len(s)

	fmt.Println(n)

	//for i := 0; i < len(s); i++ {
	//	fmt.Println(s[i])
	//	//fmt.Printf("%c\n",s[i]) //打印字符
	//}

	//字符串修改
	s2 := "白萝卜"
	s3 := []rune(s2) //把字符串转换成rune切片
	s3[0] = '红'
	fmt.Println(string(s3)) //将rune转化为 string字符串

	c1 := "红"
	c2 := '红'
	fmt.Printf("c1:%T,c2:%T\n", c1, c2)

	//类型转换
	n1 := 10
	var f float64
	f = float64(n1)
	fmt.Printf("%T\n", f) //强制转换
}
