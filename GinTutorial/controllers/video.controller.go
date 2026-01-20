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

func (c controller) Save(ctx *gin.Context){
	var video entity.Video
	err := ctx.ShouldBindJSON(&video)
	if err !=nil{
		return err;
	}
	c.service.Save(video)
	return video
}

func (c controller) FindAll() []entity.Video {
	return c.service.FindAll()
}