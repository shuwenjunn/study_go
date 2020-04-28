package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	resp, err := http.Get("http://127.0.0.1:9000/xxx/?name=哈哈&age=18")
	if err != nil {
		fmt.Println(err)
	}
	//冲resp中读取数据
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(b))
}
