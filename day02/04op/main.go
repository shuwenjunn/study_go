package main

import "fmt"

func main() {
	var (
		a = 5
		b = 2
	)

	//算数运算符
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b)
	a++ //单独的语句不能放在=后面赋值
	b--
	fmt.Println(a, b)

	//关系运算符
	fmt.Println(a == b) //go语言是强类型，只有相同的类型才能够比较
	fmt.Println(a != b) //go语言是强类型，只有相同的类型才能够比较
	fmt.Println(a >= b) //go语言是强类型，只有相同的类型才能够比较
	fmt.Println(a <= b) //go语言是强类型，只有相同的类型才能够比较
	fmt.Println(a < b)  //go语言是强类型，只有相同的类型才能够比较
	fmt.Println(a > b)  //go语言是强类型，只有相同的类型才能够比较

	//逻辑运算符 同js && || ！ 分别代表与或非

	//age := 22
	//if age > 18 && age < 60 {
	//	fmt.Println("苦逼上班族")
	//} else {
	//	fmt.Println("不用上班")
	//}

	//	位运算符 针对二进制
	// 5 的二进制 101
	// 2 的二进制  10
	//& 按位与

	fmt.Println(5 & 2) // 000

	// 按位或者 两个有一个为1就为1
	fmt.Println(5 | 2) //111 -》十进制 输出7

	//^ 异或 两位不一样就为1
	fmt.Println(5 ^ 2) //111

	//<< 将二进制位移动位数
	fmt.Println(5 << 1) // 101向左移动1位变成 1010=》转为二进制就是10

	// >> 向右移动指定位数
	fmt.Println(5 >> 1) //101 往右移一位变成 10
	fmt.Println(5 >> 3) //101 往右移3位变成 0

	//var m = int(8) //int8 类型只有8位 移动10位则位数不够。所以报错
	//fmt.Println(m << 10)

	//赋值运算符
	var x int
	x = 10
	x += 1
	x -= 1
	x *= 2
	x /= 2
	x &= 2
	x |= 2
	x ^= 2
	x <<= 2
	x >>= 2
	fmt.Println(x)

}
