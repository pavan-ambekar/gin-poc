package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pavan-ambekar/gin-poc/controller"
	"github.com/pavan-ambekar/gin-poc/middlewares"
	"github.com/pavan-ambekar/gin-poc/service"
	gindump "github.com/tpkeeper/gin-dump"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func main() {
	server := gin.New()

	server.Use(gin.Recovery(), middlewares.Logger(), middlewares.BasicAuth(), gindump.Dump())

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	server.GET("/videos", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, videoController.FindAll())
	})

	server.POST("/videos", func(ctx *gin.Context) {
		err := videoController.Save(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusCreated, gin.H{"message": "video created"})
		}
	})

	server.Run(":8080")
}
