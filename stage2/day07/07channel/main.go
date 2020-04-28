package main

import (
	"fmt"
	"sync"
)

var ch1 chan int // 通道是引用类型 穿的就是地址

var wg sync.WaitGroup
var once sync.Once

//生产100个数放到ch1中
func f1(ch1 chan<- int) { //单向通道
	defer wg.Done()
	for i := 0; i < 10; i++ {
		ch1 <- i
	}
	close(ch1)
}

func f2(ch1 <-chan int, ch2 chan<- int, s string) {
	defer wg.Done()
	for {
		x, ok := <-ch1 //从通道1里面取值 ，直到取完之后就退出循环
		fmt.Println(s, s, s, s, s, s, x)
		if !ok {
			break
		}
		ch2 <- x * x // 将通道1 里面的值放到通道2 中

	}
	once.Do(func() {
		close(ch2)
	})
}

func main() {
	a := make(chan int, 10)
	b := make(chan int, 10)
	wg.Add(2)
	go f1(a)
	go f2(a, b, "*")
	go f2(a, b, "#")

	wg.Wait()
	for ret := range b {
		fmt.Println(ret)
	}
}
