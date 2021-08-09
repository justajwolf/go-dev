package main

import (
	"github.com/changbaihe/go-manual/go-basic/attempt"
	"github.com/changbaihe/go-manual/web-service-gin/albums"
	"github.com/gin-gonic/gin"
)

func main() {
	// gin.SetMode(gin.ReleaseMode)
	attempt.PrintOne()
	router := gin.Default()
	router.GET("/albums", albums.GetAlbums)
	router.POST("/albums", albums.PostAlbums)
	router.GET("/albums/:id", albums.GetAlbumByID)
	router.Run()
}
