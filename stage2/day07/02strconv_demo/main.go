package main

import (
	"fmt"
	"strconv"
)

func main() {
	//字符串转数字
	a := "10000"
	ret1, _ := strconv.ParseInt(a, 10, 64)
	fmt.Printf("%#v %T\n", ret1, ret1)

	// 字符串 to 数字
	ret2, _ := strconv.Atoi(a)
	fmt.Printf("%#v\n", ret2)

	i := 97
	ret3 := strconv.Itoa(i)
	fmt.Println(string(i), ret3)
	fmt.Printf("%#v %T\n", ret3, ret3)

	boolStr := "true"
	val, _ := strconv.ParseBool(boolStr)
	fmt.Printf("%#v %T\n", val, val)

	//从字符串解析出浮点
	floatStr := "123.43"
	val1, _ := strconv.ParseFloat(floatStr, 64)
	fmt.Printf("%#v %T\n", val1, val1)
}
