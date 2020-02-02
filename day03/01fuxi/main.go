package main

import "fmt"

func main() {
	// array
	var ages [30]int
	fmt.Println(ages)

	// 多维数组只有最外层能够使用 ... 表示类型
	var a1 = [...][2]int{
		[2]int{1, 3},
		[2]int{3, 4},
	}
	fmt.Println(a1)

	// 数组是值类型
	x := [3]int{1, 2, 3}
	f1(x)
	fmt.Println(x)

	// 切片的定义 只能存放相同类型的元素 引用类型声明之后需要初始化
	var s1 []int // 只声明但是没有分配内存 ==nil
	// 切片定义后需要初始hua
	s1 = []int{1, 2, 3}
	fmt.Println(s1)
	// Make 初始化 分配内存
	s2 := make([]bool, 2, 4) // 三个值，type 长度和容量
	fmt.Println(s2)
	s3 := make([]int, 0, 4) // s3 不为nil 因为make已经给切片分配的内存
	fmt.Println(s3 == nil)  // 输出：false

	// 切片再切片
	// s4 := [...]int{1, 2, 3, 4, 5, 6}
	// s5 := s4[1:4]
	// s6 := s5[2:6]
	// s7 := s4[3:6]
	// fmt.Println(s4, s5, s6, s7)

	// 切片 不存值，他就像一个框 去底层数组框值

	// append 会自动扩容并且初始化

	var s8 []int

	s8 = append(s8, 1)
	fmt.Println(s8)

	//copy

	s9 := []int{1, 2, 3}
	var s10 []int
	s10 = make([]int, 3, 3)
	copy(s10, s9)
	fmt.Println(s9, s10)

	// 指针 go 只能读取指针的内存地址

	addr := "沙河"
	addrP := &addr
	fmt.Println(addr, addrP)

	fmt.Println(*addrP) // 根据内存地址查找值

	// map 存储的是 键值对的内存关系
	var m1 map[string]int
	fmt.Println(m1 == nil)
	m1 = make(map[string]int, 10) // map 在使用时需要提前预估容量
	m1["湖北"] = 18

	//判断键值 是否存在 如果返回的是bool值，我们通常用ok来接收
	value, ok := m1["姬无命"]
	fmt.Println(value, ok)
	delete(m1, "理想") // 删除的key 不存在，则什么都不做
	delete(m1, "湖北") // 删除key值为 “湖北”的一项
}

func f1(a [3]int) {

	// go 语言中函数传递的都是值
	a[1] = 100 // 此处修改的是副本的值
}
