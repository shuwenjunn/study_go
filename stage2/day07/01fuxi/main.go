package main

import (
	"encoding/json"
	"fmt"
)

type student struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

//反射
func main() {
	str := `{"name":"shu","age":9000}`
	var stu1 student
	json.Unmarshal([]byte(str), &stu1)
	fmt.Println("%#v\n", stu1)
}
