package main

import "fmt"

func main() {
	ch1 := make(chan int, 2)

	for i := 0; i < 10; i++ {
		select {
		case ch1 <- i:
			fmt.Println("存")
		case <-ch1:
			fmt.Println("取")
		}
	}
}
