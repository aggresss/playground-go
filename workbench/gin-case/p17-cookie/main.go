package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/cookie", func(ctx *gin.Context) {
		cookie, err := ctx.Cookie("gin_cookie")
		if err != nil {
			cookie = "NotSet"
			ctx.SetCookie("gin_cookie", "test", 3600, "/", "localhost", false, true)
			return
		}
		fmt.Printf("Cookie value: %s\n", cookie)
	})

	router.Run(":8080")
}
