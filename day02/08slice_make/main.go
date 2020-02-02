package main

import "fmt"

func main() {
	// make 函数创建切片
	s1 := make([]int, 5, 10) // 长度为5 容量为10
	fmt.Println(len(s1), cap(s1), s1)

	s2 := make([]int, 0, 10) // 长度为0 容量为10
	fmt.Println(len(s2), cap(s2), s2)

	//切片的赋值
	s3 := []int{1, 3, 5}
	s4 := s3 //引用类型 s3 和 s4 指向了同一个底层数组
	s3[0] = 1000
	fmt.Println(s3, s4) //引用类型

	//切片的遍历
	for i := 0; i < len(s3); i++ {
		fmt.Println(s3[i])
	}
}
