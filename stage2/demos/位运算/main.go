package main

import (
	. "fmt"
)

func main() {
	const a = 2  //0010
	const b = 6  //0110
	const c = 11 //1011

	var d = a ^ c //0010 1011 ->1001
	Println(d)    //9

	d = b & c  // 0110  1011 -> 0010
	Println(d) // 2

	var a1 = a | c //0010 1011 -> 1011
	println(a1)    //11

	var a2 = a << 2 //0010 -> 1000
	println(a2)     //8

	var a3 = b >> 2 //0110 -> 0001
	println(a3)     //1
}
