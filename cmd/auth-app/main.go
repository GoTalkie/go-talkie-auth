package main

import (
	"github.com/GoTalkie/go-talkie-auth/internal/handlers"
	"github.com/GoTalkie/go-talkie-auth/internal/middlewares"
	"github.com/GoTalkie/go-talkie-auth/internal/models"
	"github.com/gin-gonic/gin"
)

func main() {
	models.ConnectDataBase()

	r := gin.Default()
	r.Use(middlewares.RequestLogger())
	public := r.Group("/api")

	conf := handlers.Config{}

	public.POST("/register", conf.Register)
	public.POST("/login", conf.Login)

	protected := r.Group("/api/admin")
	protected.Use(middlewares.JwtAuthMiddleware())
	protected.GET("/validate", conf.User)

	r.Run(":80")
}
