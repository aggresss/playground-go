package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/JSONP", func(ctx *gin.Context) {
		data := map[string]interface{}{
			"foo": "bar",
		}

		ctx.JSONP(http.StatusOK, data)
	})

	r.Run()
}

// curl -X GET -v localhost:8080/JSONP?callback=x
