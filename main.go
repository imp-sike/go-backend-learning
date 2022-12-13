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

	r.POST("/post", controllers.PostCreate)
	r.GET("/post", controllers.PostsIndex)
	r.GET("/post/:id", controllers.PostsSingle)
	r.PUT("/post/:id", controllers.PostUpdate)
	r.DELETE("/post/:id", controllers.PostDelete)

	r.Run()
}
