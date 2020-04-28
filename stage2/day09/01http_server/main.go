package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func f1(w http.ResponseWriter, r *http.Request) {
	str := "hello world"
	w.Write([]byte(str))
}

func f2(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Query().Get("name"))
	fmt.Println(r.URL.Query().Get("age"))
	fmt.Println(r.Method)
	fmt.Println(ioutil.ReadAll(r.Body))
	w.Write([]byte("ok"))
}

func main() {
	http.HandleFunc("/home", f1)
	http.HandleFunc("/xxx/", f2)

	http.ListenAndServe("127.0.0.1:9000", nil)
}
