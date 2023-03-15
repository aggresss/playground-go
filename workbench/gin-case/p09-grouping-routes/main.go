package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func loginEndpoint(ctx *gin.Context) {
	ctx.String(http.StatusOK, "login")
}

func submitEndpoint(ctx *gin.Context) {
	ctx.String(http.StatusOK, "submit")
}

func readEndpoint(ctx *gin.Context) {
	ctx.String(http.StatusOK, "read")
}

func main() {
	router := gin.Default()

	// Simple group: v1
	v1 := router.Group("/v1")
	{
		v1.POST("/login", loginEndpoint)
		v1.POST("/submit", submitEndpoint)
		v1.POST("/read", readEndpoint)
	}

	// Simple group: v2
	v2 := router.Group("/v2")
	{
		v2.POST("/login", loginEndpoint)
		v2.POST("/submit", submitEndpoint)
		v2.POST("/read", readEndpoint)
	}

	router.Run(":8080")
}
