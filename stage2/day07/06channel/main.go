package main

import (
	"fmt"
	"sync"
)

var i []int

var b chan int // 需要指定通道中的值得类型 引用类型

var wg sync.WaitGroup

func noBuffChannel() {
	b = make(chan int) //  通道的初始化 必须使用make初始化才能使用 slice map channel
	wg.Add(1)
	go func() {
		defer wg.Done()
		x := <-b
		fmt.Println("我是goroutine 里的", x)
	}()
	b <- 10
	wg.Wait()
}

func buffChannel() {
	b = make(chan int, 1) //  通道的初始化 必须使用make初始化才能使用 slice map channel
	b <- 10
	fmt.Println("将10发送到通道中了")
	b <- 10
	fmt.Println("将20发送到通道中了")
	x := <-b
	fmt.Println("xxx", x)
}

func main() {
	//noBuffChannel()
	buffChannel()
}
