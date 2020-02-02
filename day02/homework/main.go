package main

import (
	"fmt"
	"strings"
)

func main() {
	var s = "how do you do do"
	var slice = strings.Split(s, " ")
	var l = len(slice)
	var m = make(map[string]int, l) //初始化map
	for i, v := range slice {
		_, ok := m[v]
		if !ok {
			times := 0
			for j := i; j < l; j++ {
				if slice[j] == v {
					times++
				}
			}
			m[v] = times
		}
	}

	fmt.Println(m)
}
