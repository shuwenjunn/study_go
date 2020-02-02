package main

import "fmt"

//append 为切片添加元素
func main() {
	s1 := []string{"北京", "上海", "广东"}
	//s1[3] = "广州" // 错误的写法 编译错误 索引越界
	//fmt.Println(s1)
	fmt.Println(len(s1), cap(s1))
	s1 = append(s1, "广州", "河南", "湖北", "湖南") //调用append 函数扩容必须用原来的切片接受返回值
	fmt.Println(len(s1), cap(s1), s1)

	ss := []string{"武汉", "南京", "无锡"}

	s1 = append(s1, ss...) // ...表示拆开
	fmt.Println(s1,ss)
}
