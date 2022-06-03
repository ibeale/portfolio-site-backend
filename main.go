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
	router.POST("/posts", createPost)
	router.Run(":8080")
}

func getPosts(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, posts)
}

func createPost(c *gin.Context) {
	var post models.Post
	if err := c.BindJSON(&post); err != nil {
		return
	}
	posts = append(posts, post)
	c.JSON(http.StatusCreated, post)
}
