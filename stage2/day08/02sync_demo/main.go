package main

import (
	"fmt"
	"sync"
	"time"
)

//读写互斥锁

var x = 0
var lock sync.Mutex
var wg sync.WaitGroup
var rwLock sync.RWMutex

func read() {
	defer wg.Done()
	rwLock.RLock()
	fmt.Println(x)
	time.Sleep(time.Millisecond)
	rwLock.RUnlock()
}

func write() {
	defer wg.Done()
	rwLock.Lock()
	x++
	time.Sleep(5 * time.Millisecond)
	rwLock.Unlock()
}

func main() {
	start := time.Now()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go write()
	}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go read()
	}
	wg.Wait()
	end := time.Now()
	fmt.Println(end.Sub(start))
}
