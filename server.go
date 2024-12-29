package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gitlab.com/ParkJong-Hun/gin-tutorial/controller"
	"gitlab.com/ParkJong-Hun/gin-tutorial/middlewares"
	"gitlab.com/ParkJong-Hun/gin-tutorial/service"
	"gitlab.com/ParkJong-Hun/gin-tutorial/utility"
	"io"
	"os"
	"path/filepath"
	"time"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
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

	server.Use(
		gin.Recovery(),
		middlewares.Logger(),
		middlewares.BasicAuth(),
	)

	server.GET("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.FindAll())
	})

	server.POST("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.Save(ctx))
	})

	server.Run(":8080")
}
