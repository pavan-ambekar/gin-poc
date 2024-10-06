package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pavan-ambekar/gin-poc/controller"
	"github.com/pavan-ambekar/gin-poc/service"
)

var (
	videoService    service.VideoService       = service.New()
	VideoController controller.VideoController = controller.New(videoService)
)

func main() {
	server := gin.Default()

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	server.GET("/videos", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, VideoController.FindAll())
	})

	server.POST("/videos", func(ctx *gin.Context) {
		ctx.JSON(http.StatusCreated, VideoController.Save(ctx))
	})

	server.Run(":8080")
}
