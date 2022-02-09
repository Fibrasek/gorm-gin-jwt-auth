package main

import (
	"github.com/gin-gonic/gin"

	"github.com/fibrasek/bookstore/controllers"
	"github.com/fibrasek/bookstore/middlewares"
	"github.com/fibrasek/bookstore/models"
)

func main() {
	models.ConnectDatabase()

	r := gin.Default()

	public := r.Group("/api")
	public.POST("/register", controllers.Register)
	public.POST("/login", controllers.Login)

	protected := r.Group("/api/admin")
	protected.Use(middlewares.JwtAuthMiddleware())
	protected.GET("/user", controllers.CurrentUser)

	r.GET("/books", controllers.FindBooks)
	r.POST("/books", controllers.CreateBook)
	r.GET("/books/:id", controllers.FindBook)
	r.PATCH("/books/:id", controllers.UpdateBook)
	r.DELETE("/books/:id", controllers.DeleteBook)

	r.Run()
}
