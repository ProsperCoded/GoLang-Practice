package controllers

import (
	"example.com/m/GinTutorial/entity"
	"example.com/m/GinTutorial/service"
	"example.com/m/GinTutorial/validators"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type VideoController interface {
	Save(ctx *gin.Context) (entity.Video, error)
	FindAll() []entity.Video
	ShowAll(ctx *gin.Context)
} 

type controller struct {
	service service.VideoService
}

var validate *validator.Validate
func New(videoService service.VideoService) VideoController {

	validate = validator.New()
	validate.RegisterValidation("is-cool", validators.IsCool)
	return controller{service: videoService}
}

func (c controller) Save(ctx *gin.Context) (entity.Video, error) {
	var video entity.Video
	err := ctx.ShouldBindJSON(&video)
	if err != nil {
		return entity.Video{}, err
	}

	// for custom validation
	err = validate.Struct(video)
	if err != nil {
		return entity.Video{}, err
	}
	c.service.Save(video)
	return video, nil
}

func (c controller) FindAll() []entity.Video {
	return c.service.FindAll()
}

func (c controller) ShowAll(ctx *gin.Context) {
	videos := c.service.FindAll()

	data := gin.H{
		"Title":  "Gin Template Rendering - Video List",
		"videos": videos,
		"Time":   ctx.GetString("time"),
	}
	ctx.HTML(200, "index.html", data)
}