package controller

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/ParkJong-Hun/gin-tutorial/entity"
	"gitlab.com/ParkJong-Hun/gin-tutorial/service"
)

type VideoController interface {
	FindAll() []entity.Video
	Save(ctx *gin.Context) entity.Video
}

type controller struct {
	service service.VideoService
}

func New(service service.VideoService) VideoController {
	return &controller{
		service: service,
	}
}

func (controller *controller) FindAll() []entity.Video {
	return controller.service.FindAll()
}

func (controller *controller) Save(ctx *gin.Context) entity.Video {
	var video entity.Video
	ctx.BindJSON(&video)
	controller.service.Save(video)
	return video
}
