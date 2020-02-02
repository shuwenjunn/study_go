package main

import "fmt"

func main() {
	//切片的定义 是引用类型
	var s1 []int //定义一个存放int类型元素的切片
	var s2 []string
	fmt.Println(s1, s2)
	fmt.Println(s1 == nil) //true
	fmt.Println(s2 == nil) //true
	//初始化
	s1 = []int{1, 2, 3}
	s2 = []string{"山河", "城市", "哈哈"}
	fmt.Println(s1 == nil) //false
	fmt.Println(s2 == nil) //false

	//1.切片的长度和容量
	fmt.Println(len(s1), len(s2))
	fmt.Println(cap(s1), cap(s2))

	//2.由数组得到切片
	a1 := [...]int{1, 3, 5, 7, 9, 11, 13}
	s3 := a1[0:4] //[1 3 5 7] 基于一个数组切割 左闭右开 坐包含右不包含
	fmt.Println(s3, a1)
	s4 := a1[1:6]
	fmt.Println(s4)

	s5 := a1[:4] //相当于从0 切到4
	s6 := a1[3:] // 相当于从3 切到结束
	s7 := a1[:] // 从开始切片刀结尾

	fmt.Println(s5, s6, s7)

	fmt.Printf("%T***%T\n", a1, s7)

	//切片的容量是指底层数组的容量
	// 切片是一个引用类型 都指向一个 底层数组，底层数组的值发生改变，则切片发生 改变

	// make 函数创建切片

}
