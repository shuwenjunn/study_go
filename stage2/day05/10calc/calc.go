package calc

import "fmt"

func init()  { //import 时自动执行
	fmt.Println("我被调用了")
}

func Add(x, y int) int {
	return x + y
}
