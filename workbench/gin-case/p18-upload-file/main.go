package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.MaxMultipartMemory = 8 << 20
	router.POST("/upload", func(ctx *gin.Context) {
		form, _ := ctx.MultipartForm()
		files := form.File["upload[]"]

		for _, file := range files {
			log.Println(file.Filename)
			ctx.SaveUploadedFile(file, "./")
		}
		ctx.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))
	})

	router.Run(":8080")
}

/*
curl --trace -X POST http://localhost:8080/upload \
  -F "upload[]=@${PWD}/1.log" \
  -F "upload[]=@${PWD}/2.log" \
  -H "Content-Type: multipart/form-data"
*/
