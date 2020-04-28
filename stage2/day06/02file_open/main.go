package main

import (
	"fmt"
	"os"
)

func main() {
	//f1()
	f2()
}

func f1() {
	fileObj, err := os.Open("./main.go") //当err 有值得时候，fileObj 为nil,nil 是不能调用fileObj.Close() 所有会panic
	if err != nil {
		fmt.Println("打开失败", err)
		return
	}
	defer fileObj.Close()
	fmt.Println(fileObj)
}

func f2() {
	fileObj, err := os.OpenFile("./sb.txt", os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("err", err)
	}
	defer fileObj.Close()
	fileObj.Seek(2, 0)
	fileObj.Write([]byte{'c'})
	var ret [1]byte
	n, err := fileObj.Read(ret[:])

	if err != nil {
		fmt.Println("read failed", err)
	}
	fmt.Println(string(ret[:n]))
}
