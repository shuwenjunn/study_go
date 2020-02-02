package main

import (
	"fmt"
	"strings"
)

//string
func main() {
	fmt.Println("\"") //转义符 转义" 双引号

	s := "im ok"

	fmt.Println(s)

	//多行字符串
	s2 := `
			111
		  	222
          	333
          	 44`
	fmt.Println(s2)

	//字符创长度

	fmt.Println(len(s2))

	//字符串长度
	name := "理想"

	word := "dsb"

	//字符串拼接
	ss := name + word
	fmt.Println(ss)

	//字符串拼接
	ss1 := fmt.Sprintf("%s%s", name, word)
	fmt.Println(ss1)

	//分隔

	s3 := "1.234"
	ret := strings.Split(s3, ".")
	fmt.Println(ret)

	//包含

	fmt.Println(strings.Contains(s3, "1")) //true s3 是否包含字符串1

	//前缀和后缀
	fmt.Println(strings.HasPrefix(s3, "1")) //是否有前缀
	fmt.Println(strings.HasSuffix(s3, "4")) //是否有后缀

	s4 := "abcdea"
	fmt.Println(strings.Index(s4, "a"))     //第一个出现 a 的下标
	fmt.Println(strings.LastIndex(s4, "a")) //最后一个出现a 的下标

	//拼接
	fmt.Println(strings.Join(ret, "+"))
}
