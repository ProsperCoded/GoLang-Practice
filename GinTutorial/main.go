package main

import (
	"fmt"
	"io"
	"os"

	"example.com/m/GinTutorial/controllers"
	"example.com/m/GinTutorial/middlewares"
	"example.com/m/GinTutorial/service"
	"github.com/gin-gonic/gin"
	gindump "github.com/tpkeeper/gin-dump"
)

var (
	videoService service.VideoService = service.New()
	videoController controllers.VideoController = controllers.New(videoService)
)
func setupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}
func main() {
	setupLogOutput()
	r := gin.New()

	
	r.Use(gin.Recovery(), middlewares.Logger(), middlewares.BasicAuth())

	r.GET("/", func(c *gin.Context) {
		c.String(200, "Welcome to Gin Tutorial")
	})
	r.GET("/posts", func(c *gin.Context) {
		c.JSON(200, videoController.FindAll())
	})

	r.POST("/posts", func (ctx *gin.Context){
		ctx.JSON(200, videoController.Save(ctx))
	})
	fmt.Println("Server running at http://localhost:8080/")
	r.Run(":8080")			

}