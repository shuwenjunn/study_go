package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string,age int) person {
	return person{
		name: name,
		age:  age,
	}
}

func main() {
	var x = int32(10)

	var y int32 = 1000
	fmt.Println(x, y)

	var p person = person{
		name: "hahh ",
		age:  0,
	}
	fmt.Println(p)

	var p1 = person{
		"didi",
		18,
	}

	fmt.Println(p1)
}
