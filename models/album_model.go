package models

import (
	"errors"
	"github.com/google/uuid"
)

type Album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float32 `json:"price"`
}

var Albums = []Album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 32.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry", Price: 17.99},
}

func GetAllAlbums() []Album {
	return Albums
}

func CreateAlbum(album Album) (Album, error) {
	if album.Title == "" {
		return Album{}, errors.New("title is required")
	}

	if album.Artist == "" {
		return Album{}, errors.New("artist is required")
	}

	if album.Price <= 0 {
		return Album{}, errors.New("price must be greater than zero")
	}

	// if ID is empty, generate one
	if album.ID == "" {
		album.ID = uuid.New().String()
	}

	Albums = append(Albums, album)
	return album, nil
}

func GetAlbumByID(id string) (Album, bool) {
	for _, album := range Albums {
		if album.ID == id {
			return album, true
		}
	}
	return Album{}, false
}

func DeleteAlbumByID(id string) (bool, error) {
	for i, album := range Albums {
		if album.ID == id {
			// Eliminar el álbum usando append (antes y después del índice)
			Albums = append(Albums[:i], Albums[i+1:]...)
			return true, nil
		}
	}
	return false, errors.New("album not found")
}

func UpdateAlbumByID(id string, updated Album) (Album, error) {
	for i, album := range Albums {
		if album.ID == id {
			// Mantener el ID original, actualizar los demás campos
			updated.ID = id
			Albums[i] = updated
			return updated, nil
		}
	}
	return Album{}, errors.New("album not found")
}
