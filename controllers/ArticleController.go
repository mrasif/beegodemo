package controllers

import (
	"beegodemo/models"
	"beegodemo/repo"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Index : Get all articles
func Index(c *gin.Context) {
	articles := repo.All()
	// fmt.Println(*articles)
	c.JSON(200, gin.H{
		"data": articles,
	})
}

// Create : new article
func Create(c *gin.Context) {
	article := models.Article{}
	c.BindJSON(&article)
	id := repo.Save(article)
	article = repo.Get(id)
	c.JSON(200, gin.H{
		"data": article,
	})
}

// Show : detail routes
func Show(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	article := repo.Get(id)
	c.JSON(200, gin.H{
		"data": article,
	})
}

// Destroy : detail routes
func Destroy(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	repo.Delete(id)
	c.JSON(200, gin.H{
		"message": "Article deleted!",
	})
}
