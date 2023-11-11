package controllers

import (
	"github.com/DebPaine/go-crud/initializers"
	"github.com/DebPaine/go-crud/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Post struct {
	Title string `json:"title" binding:"required"` // binding: required means that this key is required in the request
	Body  string `json:"body" binding:"required"`
}

func CreatePost(c *gin.Context) {
	// Get data from request body
	var body Post
	// binds the incoming request body with the struct and validates it
	err := c.Bind(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create a post using the model
	post := models.Post{
		Title: body.Title,
		Body:  body.Body,
	}
	result := initializers.DB.Create(&post)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error})
		return
	}

	// Return the post as response
	c.JSON(http.StatusOK, gin.H{
		"post": post,
	})
}

func GetPosts(c *gin.Context) {
	// Get all posts
	var posts []models.Post
	result := initializers.DB.Find(&posts)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error})
		return
	}

	// Return the posts
	c.JSON(http.StatusOK, gin.H{
		"posts": posts,
	})
}

func GetPost(c *gin.Context) {
	// Get the single post by using the id from url param
	id := c.Param("id")

	// Find the post
	var post models.Post
	result := initializers.DB.First(&post, id)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error})
		return
	}

	// Return the post
	c.JSON(http.StatusOK, gin.H{
		"post": post,
	})
}

func UpdatePost(c *gin.Context) {
	// Get post id from the url param
	id := c.Param("id")

	// Get data from request body and validate it
	var body Post
	err := c.Bind(&body); if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	// Find the post
	var post models.Post
	result := initializers.DB.First(&post, id)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error})
		return
	}

	// Update the post
	updateResult := initializers.DB.Model(&post).Updates(models.Post{Title: body.Title, Body: body.Body})
	if updateResult.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": updateResult.Error})
		return
	}

	// Return the updated post
	c.JSON(http.StatusOK, gin.H{
		"updatedPost": post,
	})
}
