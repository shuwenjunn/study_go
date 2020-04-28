package main

import (
	"fmt"
	"path"
	"runtime"
)

func getInfo() {
	pc, file, line, ok := runtime.Caller(1) //0 代表当前调用的层级
	if !ok {
		fmt.Println("err")
	}
	funcName := runtime.FuncForPC(pc).Name()
	fmt.Println(funcName)
	fmt.Println(path.Base(file))
	fmt.Println(line)
}

func main() {
	getInfo()
}
