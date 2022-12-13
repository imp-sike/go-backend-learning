package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/imp-sike/gocrud/initializers"
	"github.com/imp-sike/gocrud/models"
)

func PostCreate(ctx *gin.Context) {

	// get data of post request
	var body struct {
		Body  string
		Title string
	}

	ctx.Bind(&body)

	// create a post
	post := models.Post{Title: body.Title, Body: body.Body}
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		ctx.Status(400)
		return
	}

	// return it
	ctx.JSON(
		200, gin.H{
			"message": post,
		})
}

func PostsIndex(ctx *gin.Context) {
	// Get the post
	var posts []models.Post
	initializers.DB.Find(&posts)

	// return the posts
	ctx.JSON(200, gin.H{
		"posts": posts,
	})
}

func PostsSingle(ctx *gin.Context) {
	// Get the id
	id := ctx.Param("id")

	// Get the post
	var post models.Post
	initializers.DB.First(&post, id)

	// return the posts
	ctx.JSON(200, gin.H{
		"post": post,
	})
}

func PostUpdate(ctx *gin.Context) {
	// Get the id from the url
	id := ctx.Param("id")

	// get the data off req body
	var body struct {
		Title string
		Body  string
	}

	ctx.Bind(&body)

	// find the post we are updating
	var post models.Post
	initializers.DB.First(&post, id)

	// update it
	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body:  body.Body,
	})

	// respond it
	ctx.JSON(200, gin.H{
		"post": post,
	})
}

func PostDelete(ctx *gin.Context) {
	// get the id
	id := ctx.Param("id")

	// Delete the post
	initializers.DB.Delete(&models.Post{}, id)

	// respond
	ctx.Status(200)
}
