package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func f() {
	rand.Seed(time.Now().UnixNano()) //保证每次执行的时候都不一样 加随机数种子
	for i := 0; i < 5; i++ {
		r1 := rand.Int()
		r2 := rand.Intn(10) //
		fmt.Println(r1, r2)
	}

}

func f1(i int) {
	defer wg.Done()
	fmt.Println(i)
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(3)))
}

var wg sync.WaitGroup

//程序启动后  会创建一个 main goroutine 主gorountine 去执行
func main() {
	//开启一个单独的goroutine 去 执行任务
	//f()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go f1(i)
	}
	// 如何知道这是个goroutine都结束了
	wg.Wait() //等待wg的计数器降为0
}
