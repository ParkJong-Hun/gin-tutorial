package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gitlab.com/ParkJong-Hun/gin-tutorial/controller"
	"gitlab.com/ParkJong-Hun/gin-tutorial/middlewares"
	"gitlab.com/ParkJong-Hun/gin-tutorial/service"
	"gitlab.com/ParkJong-Hun/gin-tutorial/utility"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

var (
	videoService    = service.New()
	videoController = controller.New(videoService)
)

func setupLogOutPut() {
	logDir := "./log"
	utility.MkDirIfNeeded(logDir)
	logFileName := time.Now().Format("20060102_150405")
	logFilePath := filepath.Join(logDir, logFileName)
	logFile, err := os.Create(logFilePath)
	if err != nil {
		fmt.Println("create LogFile is failed.")
		panic(err)
	}
	gin.DefaultWriter = io.MultiWriter(logFile, os.Stdout)
}

func main() {
	setupLogOutPut()

	server := gin.New()

	server.Static("/css", "./templates/css")

	server.LoadHTMLGlob("templates/*.html")

	server.Use(
		gin.Recovery(),
		middlewares.Logger(),
		middlewares.BasicAuth(),
	)

	apiRoutes := server.Group("/api")
	{

		apiRoutes.GET("/videos", func(ctx *gin.Context) {
			ctx.JSON(200, videoController.FindAll())
		})

		apiRoutes.POST("/videos", func(ctx *gin.Context) {
			err := videoController.Save(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "Video Input is Valid"})
			}
			ctx.JSON(200, videoController.Save(ctx))
		})
	}

	viewRoutes := server.Group("/view")
	{
		viewRoutes.GET("/videos", videoController.ShowAll)
	}

	server.Run(":8080")
}
