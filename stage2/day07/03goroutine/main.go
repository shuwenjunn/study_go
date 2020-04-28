package main

import (
	"fmt"
	"time"
)

func hello(i int) {
	fmt.Println("hello", i)
}

//程序启动后  会创建一个 main goroutine 主gorountine 去执行
func main() {
	//开启一个单独的goroutine 去 执行任务

	for i := 0; i < 1000; i++ {
		go func(i int) {
			fmt.Println("hello", i)
		}(i)
	}
	fmt.Println("main")
	time.Sleep(time.Second * 1)
	// main 函数结束了 由main 函数启动的 gorountine 也结束了。
}
