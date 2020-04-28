package main

// 导入系统包
import (
	"flag"
	"fmt"
)

// 定义命令行参数
var mode = flag.String("mode", "", "process mode")
//返回的是一个内存地址

func main() {
	fmt.Printf("type of mode %T,value of mode %v \n", mode, mode)
	// 解析命令行参数
	flag.Parse()

	// 输出命令行参数
	fmt.Println(*mode) //通过该内存地址取到值.
}
