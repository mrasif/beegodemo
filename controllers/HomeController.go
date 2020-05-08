package controllers

import (
	"io/ioutil"

	"github.com/gin-gonic/gin"
	"github.com/russross/blackfriday"
)

// Root : this root routing
func Root(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Welcome to article server.",
	})
}

// Docs : get markdown documentation
func Docs(c *gin.Context) {
	f, _ := ioutil.ReadFile("README.md")
	markdown := blackfriday.MarkdownCommon(f)
	c.Data(200, "text/html; charset=utf-8", markdown)
}
