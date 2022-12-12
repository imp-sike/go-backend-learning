package main

import (
	"github.com/gin-gonic/gin"
	"github.com/imp-sike/gocrud/controllers"
	"github.com/imp-sike/gocrud/initializers"
)

// init function runs before main
func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
}

func main() {
	r := gin.Default()

	r.GET("/", controllers.PostCreate)

	r.Run()
}
