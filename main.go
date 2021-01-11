package main

import (
	"github.com/gin-gonic/gin"
	
	"github.com/gerrr/blog_backend/blog_backend/controllers"
	"github.com/gerrr/blog_backend/blog_backend/models"
)

func main() {
	r := gin.Default()
	
	//Trying to connect the DB
	models.ConnectDataBase()
	
	//Routes
	r.GET("/books", controllers.FindBooks)
	r.POST("/books", controllers.CreateBook)
	r.GET("/books/:id", controllers.FindBook)
	r.PATCH("/books/:id", controllers.UpdateBook)
	r.DELETE("/books/:id", controllers.DeleteBook)
	
	//Starting API server
	r.Run()
}