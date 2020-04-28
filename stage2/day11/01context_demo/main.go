package main

import (
	"fmt"
	"sync"
	"time"
)

var exitChan chan bool = make(chan bool, 1)
var wg sync.WaitGroup

//为什么需要context

func f() {
	defer wg.Done()
FORLOOP:
	for {
		fmt.Println("周玲")
		time.Sleep(time.Millisecond * 500)
		select {
		case <-exitChan:
			break FORLOOP
		default:

		}
	}
}

func main() {
	wg.Add(1)
	go f()
	time.Sleep(time.Second * 5)
	exitChan <- true
	wg.Wait()
}
