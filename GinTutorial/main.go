package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"example.com/m/GinTutorial/controllers"
	"example.com/m/GinTutorial/middlewares"
	"example.com/m/GinTutorial/service"
	"github.com/gin-gonic/gin"
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

	// Load templates
	r.LoadHTMLGlob("templates/*")

	apiRoutes := r.Group("/api")
	apiRoutes.Use(gin.Recovery(), middlewares.Logger(), middlewares.BasicAuth())

	apiRoutes.GET("/", func(c *gin.Context) {
		c.String(200, "Welcome to Gin Tutorial")
	})
	apiRoutes.GET("/posts", func(c *gin.Context) {
		c.JSON(200, videoController.FindAll())
	})

	apiRoutes.POST("/posts", func (ctx *gin.Context){
		_, err := videoController.Save(ctx)

		if err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})	
		}else {
			ctx.JSON(http.StatusOK, gin.H{"message": "Video is valid and saved"})
		}
	})

	// HTML routes for template rendering
	r.GET("/", func(c *gin.Context) {
		videoController.ShowAll(c)
	})

	fmt.Println("Server running at http://localhost:8080/")
	r.Run(":8080")			

}