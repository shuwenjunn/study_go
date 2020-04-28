package main

import "fmt"

// 结构体 是值类型
type person struct {
	name, gender string
}

func f(x person) {
	x.gender = "女"
}

//指针修改
func f2(x *person) {
	(*x).gender = "女" // 根据内存地址找到原变量的内存地址 修改的就是原变量

}

func main() {
	var p person
	p.name = "哈哈"
	p.gender = "男"

	f2(&p)
	fmt.Println(p)

	//结构体指针1
	var p2 = new(person) // person 类型的指针
	(*p2).name = "理想"
	fmt.Printf("%T%p\n", p2, p2)
	//fmt.Println(p2)

	// 结构体 指针2 key value  初始化
	var p3 = person{
		name:   "谢谢 ",
		gender: "男",
	}
	fmt.Printf("%#v\n", p3)

	p4 := person{"小王子", "男"} // 值得顺序 要和定义结构体得顺序一致
	fmt.Println(p4)
}
