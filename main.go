package main

import (
	"baebeez-auth/controller"
	"baebeez-auth/middleware"
	"baebeez-auth/models"

	"github.com/gin-gonic/gin"
)

func main() {

	models.ConnectDataBase()

	r := gin.Default()

	public := r.Group("/auth")

	public.POST("/login", controller.Login)
	public.POST("/register", controller.Register)
	public.POST("/saveprofile", controller.SaveProfile) // will be protected after mail verification is done
	protected := r.Group("/protected")
	protected.Use(middleware.JwtAuthMiddleWare())
	protected.GET("/user", controller.CurrentUser)
	protected.POST("/logout", controller.Logout)

	r.Run(":8080")
}
