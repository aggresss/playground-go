package main

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		t := time.Now()
		ctx.Set("example", "12345")
		// before request
		ctx.Next()
		// after request
		log.Print(time.Since(t))
		log.Println(ctx.Writer.Status())
	}
}

func main() {
	r := gin.New()
	r.Use(Logger())

	r.GET("/test", func(ctx *gin.Context) {
		example := ctx.MustGet("example").(string)
		log.Println(example)
	})

	r.Run()
}
