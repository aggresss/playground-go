package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var secrets = gin.H{
	"foo": gin.H{"email": "foo@bar.com", "phone": "1234"},
}

func main() {
	r := gin.Default()

	authorized := r.Group("/admin", gin.BasicAuth(gin.Accounts{
		"foo": "bar",
	}))

	authorized.GET("/secrets", func(ctx *gin.Context) {
		user := ctx.MustGet(gin.AuthUserKey).(string)
		if secret, ok := secrets[user]; ok {
			ctx.JSON(http.StatusOK, gin.H{"user": user, "secret": secret})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"user": user, "secret": "NO SECRET"})
		}
	})

	r.Run(":8080")
}

// curl http://localhost:8080/admin/secrets -u foo:bar
