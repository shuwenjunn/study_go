// 指针的语法糖 *
package main

import "fmt"

type testStruct struct {
	age  int
	name string
}

func changeAge(t *testStruct) {
	t.age++ //等价于 *t.age++,意思是取到入参指针对应的变量 然后+1 从新放到该指针中
}

func main() {
	shu := &testStruct{
		age:  100,
		name: "舒文俊",
	}
	fmt.Println(*shu)
	changeAge(shu)
	fmt.Println(*shu)
}
