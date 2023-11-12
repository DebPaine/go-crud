package main

import (
	"github.com/DebPaine/go-crud/controllers"
	"github.com/DebPaine/go-crud/initializers"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload" // Since we only want this import for it's side-effect
)

// init function is run before main
func init() {
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()
	r.POST("/post", controllers.CreatePost)
	r.GET("/posts", controllers.GetPosts)
	r.GET("/posts/:id", controllers.GetPost)
	r.PUT("/posts/:id", controllers.UpdatePost)
	r.DELETE("/posts/:id", controllers.DeletePost)
	r.Run() // This will run using PORT from .env, else :8080 by default
}
