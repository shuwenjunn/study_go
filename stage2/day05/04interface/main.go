package main

import "fmt"

type car interface {
	run()
}

type falali struct {
	brand string
}

type baoshijie struct {
	brand string
}

func drive(c car) {
	c.run()
}

func (f falali) run() {
	fmt.Println("速度70", f.brand)
}

func (b baoshijie) run() {
	fmt.Println("速度80", b.brand)
}

func main() {
	var f1 falali = falali{brand: "法拉利"}
	var p1 baoshijie = baoshijie{brand: "保时捷"}

	drive(f1)
	drive(p1)
}
