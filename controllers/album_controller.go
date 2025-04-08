package controllers

import (
	"firstWebApi/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAllAlbums(c *gin.Context) {
	albums := models.GetAllAlbums()
	c.JSON(http.StatusOK, albums)
}

func PostAlbums(c *gin.Context) {
	var newAlbum models.Album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	createdAlbum, err := models.CreateAlbum(newAlbum)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdAlbum)

}

func GetAlbumById(c *gin.Context) {
	id := c.Param("id")

	album, found := models.GetAlbumByID(id)
	if !found {
		c.JSON(http.StatusNotFound, gin.H{"message": "album not found"})
		return
	}

	c.JSON(http.StatusOK, album)
}

func DeleteAlbum(c *gin.Context) {
	id := c.Param("id")

	ok, err := models.DeleteAlbumByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if ok {
		c.JSON(http.StatusOK, gin.H{"message": "album deleted"})
	}
}

func PutAlbum(c *gin.Context) {
	id := c.Param("id")
	var updatedAlbum models.Album

	//Intenta deserializar el JSON recibido y convertirlo a un objeto Album.
	//Si el cuerpo está mal formado (por ejemplo, falta un campo o es inválido), se retorna un error.
	if err := c.ShouldBindJSON(&updatedAlbum); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	album, err := models.UpdateAlbumByID(id, updatedAlbum)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, album)
}
