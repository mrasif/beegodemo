package controllers

import "github.com/gin-gonic/gin"

// Root : this root routing
func Root(c *gin.Context){
	c.JSON(200, gin.H{
		"message": "Welcome to article server.",
	})
}