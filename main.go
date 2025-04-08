package main

import (
	"firstWebApi/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/albums", controllers.GetAllAlbums)
	router.POST("/albums", controllers.PostAlbums)
	router.GET("/albums/:id", controllers.GetAlbumById)
	router.DELETE("/albums/:id", controllers.DeleteAlbum)
	router.PUT("/albums/:id", controllers.PutAlbum)

	err := router.Run("localhost:8080")
	if err != nil {
		return
	}
}
