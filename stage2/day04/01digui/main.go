package main

import "fmt"

//递归 自己调用自己

//阶乘 5*4*3*2*1

func f(n uint64) uint64 {
	if n <= 1 {
		return 1
	}
	return n * f(n-1)
}

func main() {
	ret := f(5)
	fmt.Println(ret)

	fmt.Println(f1(3))
}

//n个台阶 一次可以走一步或者两步，一共有多少种走法

func f1(n uint64) uint64 {
	// 出口
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	return f1(n-1) + f1(n-2)
}
