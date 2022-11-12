package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
	})
	err := r.Run()
	if err != nil {
		panic(err)
	}
}
