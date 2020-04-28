package main

import (
	"fmt"
	"os"
)

func main() {
	fileObj, err := os.Open("./main.go")
	if err != nil { //err 不为空 表示有错误
		fmt.Println("open fialed", err)
		return
	}

	defer fileObj.Close()
	var tmp = make([]byte, 128)

	for {
		n, err := fileObj.Read(tmp)
		if err != nil {
			fmt.Println("读出错")
			return
		}

		fmt.Printf("%d\n", n)

		fmt.Println(string(tmp))

		if n <128 {
			return
		}
	}
}
