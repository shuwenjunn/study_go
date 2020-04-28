package main

import (
	"encoding/json"
	"fmt"
)

type people struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	s := `[{"name":"shu","age":"18"},{"name":"fad","age":22}]`

	res := make([]*people, 2)
	_ = json.Unmarshal([]byte(s), &res)
	fmt.Printf("res--- %#v \n", res)

	for _, v := range res {
		fmt.Printf("%#v \n", *v)
	}
}
