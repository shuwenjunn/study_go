package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

//使用goroutine和channel实现一个计算int64随机数各位数和的程序。
//开启一个goroutine循环生成int64类型的随机数，发送到jobChan
//开启24个goroutine从jobChan中取出随机数计算各位数的和，将结果发送到resultChan
//主goroutine从resultChan取出结果并打印到终端输出

type jobChan struct {
	x int64
}

type resultChan struct {
	jobChan
	result int64
}

var job chan jobChan
var result chan resultChan
var wg sync.WaitGroup

func createRandNum() (num int64) {
	rand.Seed(time.Now().UnixNano())
	num = rand.Int63()
	return num
}

func calcNum(num int64) (sum int64) {
	sum = 0
	for num > 0 {
		sum += num % 10 //123
		num = num / 10
	}

	return
}

func creatJob(job chan<- jobChan) {
	defer wg.Done()
	for {
		x := createRandNum()
		job <- jobChan{x: x}
		time.Sleep(time.Second)
	}
}

func worker(job <-chan jobChan, result chan<- resultChan) {
	defer wg.Done()
	jobItem := <-job
	num := calcNum(jobItem.x)
	resultItem := resultChan{
		jobChan: jobItem,
		result:  num,
	}
	result <- resultItem
}

func main() {
	job = make(chan jobChan, 100)
	result = make(chan resultChan, 100)
	wg.Add(1)
	go creatJob(job)
	wg.Add(24)
	for i := 0; i < 24; i++ {
		go worker(job, result)
	}

	for it := range result {
		fmt.Println(it.jobChan.x)
		fmt.Println(it.result)
	}
}
