package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Artist string `json:"artist"`
	Year   int    `json:"year"`
}

var albums = []album{
	{ID: "1", Title: "The Dark Side of the Moon", Artist: "Pink Floyd", Year: 1973},
	{ID: "2", Title: "Dark Side of the Moon", Artist: "Pink Floyd", Year: 1973},
	{ID: "7", Title: "The Wall", Artist: "Pink Floyd", Year: 1973},
	{ID: "8", Title: "Another Brick in the Wall, Part 1", Artist: "Pink Floyd", Year: 1973},
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func postAlbums(c *gin.Context) {
	var newAlbum album
	c.BindJSON(&newAlbum)
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, albums)
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbums)
	router.Run(":8080")
}
