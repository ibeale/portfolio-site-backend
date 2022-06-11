package main

import (
	"net/http"
	"time"

	"github.com/ibeale/portfolio-site-backend/models"

	"github.com/gin-gonic/gin"
)

var posts = []models.Post{
	{
		Title:   "Hello World",
		Content: "Welcome to my blog",
		DB_Model: models.DB_Model{
			ID:        1,
			CreatedAt: time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC),
			UpdatedAt: time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC),
			DeletedAt: nil,
		},
	},
}

func main() {
	router := gin.Default()
	models.ConnectDatabase()
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
