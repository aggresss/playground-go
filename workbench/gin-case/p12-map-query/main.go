package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/post", func(ctx *gin.Context) {
		ids := ctx.QueryMap("ids")
		names := ctx.PostFormMap("names")

		fmt.Printf("ids: %v; name: %v\n", ids, names)
	})

	router.Run()
}

/*
curl -v -X POST \
  'http://localhost:8080/post?ids\[a\]=1234&ids\[b\]=hello' \
  -H'Content-Type: application/x-www-form-urlencoded' \
  -d'names[first]=5678&names[second]=world' \
*/
