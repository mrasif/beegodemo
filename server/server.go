package server

import (
	"beegodemo/controllers"

	"github.com/gin-gonic/gin"
)

// Run : It will run server
func Run() {
	router := gin.Default()
	router.GET("/", controllers.Root)
	router.GET("/articles", controllers.Index)
	router.POST("/articles", controllers.Create)
	router.GET("/articles/:id", controllers.Show)
	router.DELETE("/articles/:id", controllers.Destroy)
	router.Run()
}
