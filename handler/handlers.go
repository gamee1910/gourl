package handler

import (
	"net/http"

	"github.com/gamee1910/gourl/shortener"
	"github.com/gamee1910/gourl/store"
	"github.com/gin-gonic/gin"
)

type UrlCreationRequest struct {
	OriginalUrl string `json:"original_url" binding:"required"`
	UserId      string `json:"user_id" binding:"required"`
}

func CreateShortUrl(c *gin.Context) {
	var creationRequest UrlCreationRequest

	if err := c.ShouldBindJSON(&creationRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	shortUrl := shortener.GenerateShortLink(creationRequest.OriginalUrl, creationRequest.UserId)

	store.SaveUrlMapping(shortUrl, creationRequest.OriginalUrl, creationRequest.UserId)

	host := "http://localhost:8080/"
	c.JSON(http.StatusOK, gin.H{
		"message":   "short url created successfully",
		"short_url": host + shortUrl,
	})

}

func HandlerShortUrlRedirect(c *gin.Context) {
	shortUrl := c.Param("shortUrl")
	initialUrl := store.RetrieveInitialUrl(shortUrl)

	c.Redirect(302, initialUrl)
}
