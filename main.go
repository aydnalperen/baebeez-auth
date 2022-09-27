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
	// will be protected after mail verification is done
	protected := r.Group("/protected")
	protected.Use(middleware.JwtAuthMiddleWare())
	protected.POST("/saveprofile", controller.SaveProfile)
	protected.GET("/user", controller.CurrentUser)
	protected.POST("/logout", controller.Logout)

	protected.GET("/getuserbyuid", controller.GetUserByUid)
	protected.GET("/getusers", controller.GetUsers)

	protected.POST("/update/", controller.UpdateUser)
	protected.POST("/images/", controller.UploadImage)

	protected.POST("/like", controller.AddToLikes)
	protected.GET("/matches", controller.GetMatches)
	protected.POST("/verify-user", controller.VerifyUser)
	protected.Static("/images", "./images")

	r.Run(":8080")
}
