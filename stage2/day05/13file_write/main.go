package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func writeDemo1() {
	//打开文件写内容
	fileObj, err := os.OpenFile("./xx.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("打开失败")
	}

	//write
	fileObj.Write([]byte("wo sha le "))
	fileObj.WriteString("解释不了了")
	fileObj.Close()
}

//bufio
func writeDemo2() {
	//打开文件写内容
	fileObj, err := os.OpenFile("./xx.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("打开失败")
	}
	defer fileObj.Close()
	//创建一个写的对象 写到缓存区里面的
	wr := bufio.NewWriter(fileObj)
	wr.WriteString("用bufio的方式写")
	wr.Flush()
}

func writeDemo3() {
	err := ioutil.WriteFile("./xx.txt", []byte("哈哈哈"), 0644)
	if err != nil {
		fmt.Println("哈哈哈哈")
	}
}

func main() {
	//writeDemo1()
	//writeDemo2()
	writeDemo3()
}
