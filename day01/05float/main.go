package main

import (
	"fmt"
	"math"
)

//浮点数
func main() {
	fmt.Println(math.MaxFloat32)
	f1 := 1.23456
	fmt.Printf("%T\n", f1) //默认go中的小数都是float64
	f2 := float32(1.23456) //强制定义32 位浮点数
	fmt.Printf("%T\n", f2)
	//f1 = f2 //float32 类型的值 不能赋值给64
}
