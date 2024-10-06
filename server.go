package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pavan-ambekar/gin-poc/controller"
	"github.com/pavan-ambekar/gin-poc/middlewares"
	"github.com/pavan-ambekar/gin-poc/service"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func main() {
	server := gin.New()

	server.Use(gin.Recovery(), gin.Logger(), middlewares.Logger())

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	server.GET("/videos", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, videoController.FindAll())
	})

	server.POST("/videos", func(ctx *gin.Context) {
		ctx.JSON(http.StatusCreated, videoController.Save(ctx))
	})

	server.Run(":8080")
}
