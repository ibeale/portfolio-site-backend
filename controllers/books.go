package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ibeale/portfolio-site-backend/models"
)

// GET /posts
// Get all posts
func FindPosts(c *gin.Context) {
	var posts []models.Post
	models.DB.Find(&posts)

	c.JSON(http.StatusOK, gin.H{"data": posts})
}
