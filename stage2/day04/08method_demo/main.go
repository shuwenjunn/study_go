package main

import "fmt"

type dog struct {
	name string
}

func newDog(name string) dog {
	return dog{
		name: name,
	}
}

//方法是作用于特定类型的函数
// 接受者表示的具体的类型，多用类型名首字母小写
func (d dog) wang() {
	fmt.Printf("%swang wang wang\n", d.name)
}

func main() {
	d1 := newDog("黑狗")
	d1.wang()
}
