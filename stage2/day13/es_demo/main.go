package main

import (
	"context"
	"fmt"
	"github.com/olivere/elastic"
)

// Student ...
type Student struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Married bool   `json:"married"`
}

func main() {
	client, err := elastic.NewClient(elastic.SetURL("http://localhost:9200"))
	if err != nil {
		// Handle error
		panic(err)
	}

	fmt.Println("connect to es success")

	p1 := Student{Name: "Rion", Age: 22, Married: false}
	// 链式操作
	put1, err := client.Index().Index("student").Type("go").BodyJson(p1).Do(context.Background())
	if err != nil {
		// Handle error
		panic(err)
	}
	fmt.Printf("Indexed student %s to index %s, type %s\n", put1.Id, put1.Index, put1.Type)
}
