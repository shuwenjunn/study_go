package main

import "fmt"

const (
	n1 = 5    //0
	n2 = iota //1
	n3        //2
	n4        //3
)

func main() {
	fmt.Println(n1, n2, n3, n4)

}
