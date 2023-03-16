package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/someDataFromReader", func(ctx *gin.Context) {
		response, err := http.Get("https://raw.githubusercontent.com/gin-gonic/logo/master/color.png")
		if err != nil || response.StatusCode != http.StatusOK {
			ctx.Status(http.StatusServiceUnavailable)
			return
		}

		extraHeaders := map[string]string{
			"Content-Disposition": `attachment; filename=gopher.png`,
		}

		ctx.DataFromReader(http.StatusOK, response.ContentLength, response.Header.Get("Content-Type"), response.Body, extraHeaders)
	})

	router.Static("fs", "../")
	router.StaticFS("fs1", http.Dir("../"))

	router.Run(":8080")
}
