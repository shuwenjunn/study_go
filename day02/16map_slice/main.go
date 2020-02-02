package main

import "fmt"


// 重点 切片和map 一定要进行初始化
//map slice
func main() {
	//元素类型为map 的切片

	var s1 = make([]map[int]string, 10, 10) // 预估切片的长度是多少
	//s1[0][100]="hahha" 这种赋值方法是错误的 没有对内部的map 进行初始化

	s1[0] = make(map[int]string, 1)
	s1[0][100] = "hahha"
	fmt.Println(s1)

	// 值切片类型的map
	var m1 = make(map[string][]int, 10)
	m1["北京"] = []int{10, 20, 30}
	fmt.Println(m1)

	//统计 how do you do 每个 单词出现的次数

}
