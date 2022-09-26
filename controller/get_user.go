package controller

import (
	"baebeez-auth/models"
	"baebeez-auth/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUserByUid(ctx *gin.Context) {
	uid, _ := utils.ExtractTokenUID(ctx)
	var user models.User

	if result := models.DB.First(&user, uid); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}
	ctx.JSON(http.StatusOK, &user)
}
