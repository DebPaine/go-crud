package main

import (
	"github.com/DebPaine/go-crud/initializers"
	"github.com/DebPaine/go-crud/models"
	_ "github.com/joho/godotenv/autoload"
)

func init() {
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}

