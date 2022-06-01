package main

import (
	"net/http"
	"time"

	"github.com/ibeale/portfolio-site-backend/models"

	"github.com/gin-gonic/gin"
)

var posts = []models.Post{
	{
		ID:      "1",
		Title:   "Hello World",
		Content: "Welcome to my blog",
		Created: time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC),
	},
}

func main() {
	router := gin.Default()
	router.GET("/posts", getPosts)
	router.Run(":8080")
}

func getPosts(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, posts)
}
