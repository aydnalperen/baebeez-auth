package controller

import (
	"baebeez-auth/models"
	"baebeez-auth/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func MakeCompleted(ctx *gin.Context) { //
	uid, _ := utils.ExtractTokenUID(ctx)

	var user models.User

	if result := models.DB.Model(&models.User{}).Where("uid=?", uid).Find(&user); result != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

}
