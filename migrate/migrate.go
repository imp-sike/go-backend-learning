package main

import (
	"github.com/imp-sike/gocrud/initializers"
	"github.com/imp-sike/gocrud/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
