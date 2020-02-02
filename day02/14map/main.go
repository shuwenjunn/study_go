package main

import "fmt"

// 首先一定要初始化 ， 引用类型

func main() {
	var m1 map[string]int
	fmt.Println(m1 == nil)        //还没初始化 没有在内存中开辟空间
	m1 = make(map[string]int, 10) //初始化时要估算号该map 容量，避免程序运行的时候再动态扩容
	m1["理想"] = 18
	m1["姬无命"] = 20
	fmt.Println(m1)

	fmt.Println(m1["理想"])

	fmt.Println(m1["aaa"])
	value, ok := m1["aaa"] // 如果不存在这个key,拿到对应类型的零值

	//约定成俗 用ok 接收返回的bool值
	if !ok {
		fmt.Println("m1中无此key值")
	} else {
		fmt.Println("m1中有此key值", value)
	}

	//map 的遍历

	for key, v := range m1 {
		fmt.Println(key, v)
	}

	//删除

	delete(m1,"姬无命")
	fmt.Println(m1)
	delete(m1,"沙河") // 删除不存在的key
}
