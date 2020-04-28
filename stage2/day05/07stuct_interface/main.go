package main

//同一个结构体实现一个接口
type animal interface {
	mover
	eater
}

type mover interface {
	move()
}

type eater interface {
	eat(string)
}

type cat struct {
	name string
	feet int8
}

func (c cat) eat(s string) {

}

func (c cat) move() {

}

func main() {
	//interfact {} 空接口没有必要

}
