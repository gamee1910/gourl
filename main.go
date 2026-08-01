package main

import (
	"fmt"

	"github.com/gamee1910/gourl/handler"
	"github.com/gamee1910/gourl/store"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Hey from GOURL!",
		})
	})

	r.POST("/create", handler.CreateShortUrl)

	r.GET("/:shortUrl", handler.HandlerShortUrlRedirect)

	store.InitializeStore()

	err := r.Run(":8080")

	if err != nil {
		panic(fmt.Sprintf("Failed to start the web server - Error: [%v]", err))
	}
}
