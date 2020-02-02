package main

import "fmt"

// array
func main() {
	// 必须指定 容量和类型
	//数组的长度是数组类型的一部分
	var a1 [3]bool // true false true
	var a2 [4]bool // true false true

	fmt.Printf("%T,%T\n", a1, a2) //a1,a2的类型分别为[3]bool,[4]bool

	//数组的初始化
	//如果不初始化 默认元素都是零值
	fmt.Println(a1, a2)
	//初始化方法1
	a1 = [3]bool{true, true, true}
	//初始化方法2 不指定长度长度 由内容决定
	a10 := [...]int{1, 2, 3, 5, 6, 8}
	fmt.Println(a10)

	//初始化方法3 根据  索引初始化
	a3 := [5]int{0: 1, 4: 5}
	fmt.Println(a3)

	//数组遍历
	citys := [...]string{"北京", "上海", "深圳"}

	//根据索引遍历
	for i := 0; i < len(citys); i++ {
		fmt.Println(citys[i])
	}

	for i, v := range citys { //i 标识索引  v则是索引对应的值
		fmt.Println(i)
		fmt.Println(v)
	}

	//多维数组 [[1 2] [3 4] [5 6]]
	var a11 [3][2]int
	a11 = [3][2]int{
		[2]int{1, 2},
		[2]int{3, 4},
		[2]int{5, 6},
	}
	fmt.Println(a11)

	//多维数组的遍历
	for _, v := range a11 {
		fmt.Println(v)
		for _, w := range v {
			fmt.Println(w)
		}
	}

	//  数值 是值类型 不是引用类型,概念同js中的值类型和 引用类型
	b1 := [3]int{1, 2, 3}
	b2 := b1
	b2[0] = 100
	fmt.Println(b2, b1)

	// 练习题
	x1 := [...]int{1, 3, 5, 7, 8}
	for i1, v1 := range x1 {
		for i2, v2 := range x1 {
			if i2 != i1 && v1+v2 == 8 {
				fmt.Println(i1, i2)
			}
		}
	}
}
