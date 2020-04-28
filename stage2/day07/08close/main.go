package main

import "fmt"

func main() {
	ch1 := make(chan int, 10)
	ch1 <- 10
	ch1 <- 20
	close(ch1)
	for x := range ch1 {
		fmt.Println(x)
	}
}
