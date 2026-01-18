package controllers

import (
	"example.com/m/GinTutorial/entity"
	"example.com/m/GinTutorial/service"
	"github.com/gin-gonic/gin"
)

type VideoController interface {
	Save(ctx *gin.Context) entity.Video
	FindAll() []entity.Video
} 

type controller struct {
	service service.VideoService
}

func New(videoService service.VideoService) VideoController {
	return controller{service: videoService}
}

func (c controller) Save(ctx *gin.Context) entity.Video {
	var video entity.Video
	ctx.ShouldBindJSON(&video)
	c.service.Save(video)
	return video
}

func (c controller) FindAll() []entity.Video {
	return c.service.FindAll()
}