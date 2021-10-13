package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hngphanminh147/gin_api/controllers"
	"github.com/hngphanminh147/gin_api/db"
)

func main() {
	r := gin.Default()

	db.SetupDatabaseConnection()

	r.GET("/posts", controllers.FindPosts)
	r.POST("/posts", controllers.CreatePost)
	r.GET("/post/:id", controllers.FindPost)
	r.PUT("/post/:id", controllers.UpdatePost)
	r.DELETE("/post/:id", controllers.DeletePost)

	r.GET("/tags", controllers.FindTags)
	r.POST("/tags", controllers.CreateTag)
	r.PUT("/tag/:id", controllers.UpdateTag)

	r.Run()
}
