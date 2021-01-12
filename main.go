package main

import (
	"github.com/gin-gonic/gin"
	
	"github.com/gerrr/blog_backend/controllers"
	"github.com/gerrr/blog_backend/models"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	
	//Trying to connect the DB
	models.ConnectDataBase()
	
	//Routes
	r.GET("/articles", controllers.AllArticles)
	r.POST("/articles", controllers.CreateArticle)
	r.GET("/article/:id", controllers.FindArticle)
	r.DELETE("/article/:id", controllers.DeleteArticle)
	
	//Starting API server
	_ = r.Run()
}