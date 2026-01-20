package main

import (
	"fmt"

	"example.com/m/GinTutorial/controllers"
	"example.com/m/GinTutorial/middlewares"
	"example.com/m/GinTutorial/service"
	"github.com/gin-gonic/gin"
)

var (
	videoService service.VideoService = service.New()
	videoController controllers.VideoController = controllers.New(videoService)
)

func main() {
	r := gin.New()

	
	r.Use(gin.Recovery(), middlewares.Logger())
	r.GET("/posts", func(c *gin.Context) {
		c.String(200, "Hello, World!")
		c.JSON(200, videoController.FindAll())
	})

	r.POST("/posts", func (ctx *gin.Context){
		ctx.JSON(200, videoController.Save(ctx))
	})
	fmt.Println("Server running at http://localhost:8080/")
	r.Run(":8080")			

}