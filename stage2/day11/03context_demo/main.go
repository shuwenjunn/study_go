package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var exitChan chan bool = make(chan bool, 1)
var wg sync.WaitGroup

//为什么需要context

func f(ctx context.Context) {
	defer wg.Done()
	wg.Add(1)
	go f2(ctx)
FORLOOP:
	for {
		fmt.Println("周玲")
		time.Sleep(time.Millisecond * 500)
		select {
		case <-ctx.Done():
			break FORLOOP
		default:

		}
	}
}

func f2(ctx context.Context) {
	defer wg.Done()
FORLOOP:
	for {
		fmt.Println("保德路")
		time.Sleep(time.Millisecond * 500)
		select {
		case <-ctx.Done():
			break FORLOOP
		default:

		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)
	go f(ctx)
	time.Sleep(time.Second * 5)
	cancel()
	wg.Wait()
}
