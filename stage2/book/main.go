package main

import (
	"github.com/gin-gonic/gin"
	"net/http"

)

func main() {

	//初始化链接

	r := gin.Default()

	r.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "hello world")
	})
	r.Run(":8888")
}
