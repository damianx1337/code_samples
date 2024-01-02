package main

import (
	"psgres/controllers"
	"psgres/models"

	"github.com/gin-gonic/gin"
)

func main() {
	models.ConnectDatabase()
	//log.Println(controllers.FindPosts)

	router := gin.Default()

  router.ForwardedByClientIP = true
  router.SetTrustedProxies([]string{"127.0.0.1"})

  router.StaticFile("/favicon.ico", "./favicon.ico")

	v1 := router.Group("/api/v1")

	//v1.POST("/posts", controllers.CreatePost)
	v1.GET("/posts", controllers.FindPosts)
	//v1.GET("/posts/:id", controllers.FindPost)
	//v1.PATCH("/posts/:id", controllers.UpdatePost)
	//v1.DELETE("/posts/:id", controllers.DeletePost)

	router.Run("0.0.0.0:8080")
}
