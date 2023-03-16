/*
- https://github.com/go-playground/validator

When using the Bind-method, Gin tries to infer the binder depending on the Content-Type header.
If you are sure what you are binding, you can use MustBindWith or ShouldBindWith .

*/

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type Login struct {
	User     string `form:"user" json:"user" xml:"user" binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

func main() {
	router := gin.Default()

	router.POST("loginJSON", func(ctx *gin.Context) {
		var json Login
		if err := ctx.ShouldBindJSON(&json); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if json.User != "manu" || json.Password != "123" {
			ctx.JSON(http.StatusOK, gin.H{
				"status": "unauthorized",
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	})

	router.POST("/loginForm", func(ctx *gin.Context) {
		var form Login

		if err := ctx.ShouldBind(&form); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		if form.User != "manu" || form.Password != "123" {
			ctx.JSON(http.StatusOK, gin.H{
				"status": "unauthorized",
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	})

	router.POST("/loginXML", func(ctx *gin.Context) {
		var xml Login
		if err := ctx.ShouldBindXML(&xml); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if xml.User != "manu" || xml.Password != "123" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	})

	router.POST("/loginMultiForm", func(ctx *gin.Context) {
		var form Login
		if err := ctx.ShouldBindWith(&form, binding.FormMultipart); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if form.User != "manu" || form.Password != "123" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	})

	router.POST("/loginQuery", func(ctx *gin.Context) {
		var query Login
		if err := ctx.ShouldBindWith(&query, binding.Query); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if query.User != "manu" || query.Password != "123" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	})

	router.Run(":8080")
}

/*
curl -v -X POST \
  http://localhost:8080/loginJSON \
  -H 'content-type: application/json' \
  -d '{ "user": "manu", "password": "123" }'
*/

/*
curl -v \
  http://localhost:8080/loginMultiForm \
  --form user=manu --form password=123
*/

/*
curl -v -X POST http://localhost:8080/loginQuery?user=manu\&password=123
*/
