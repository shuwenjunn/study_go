package main

//常量

//常量定义之后不能修改 在程序运行期间不会修改

const pi = 3.14

const (
	statusOk = 200
	notFound = 404
)

//批量声明后面没赋值，默认和上一行一致
const (
	n1 = 100
	n2
	n3
)

const (
	a1 = iota //0
	a2        //1
	a3        //2
)

func main() {

}
