package controllers

import "github.com/gin-gonic/gin"

func PostCreate(ctx *gin.Context) {
	ctx.JSON(
		200, gin.H{
			"message": "Hello",
		})
}
