package main

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/long_async", func(ctx *gin.Context) {
		cCp := ctx.Copy()
		go func() {
			time.Sleep(5 * time.Second)
			log.Println("Done! in path", cCp.Request.URL.Path)
		}()
	})

	r.GET("long_sync", func(ctx *gin.Context) {
		time.Sleep(5 * time.Second)
		log.Println("Done! in path", ctx.Request.URL.Path)
	})

	r.Run(":8080")
}
