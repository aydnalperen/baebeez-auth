package controller

import (
	"baebeez-auth/models"
	"baebeez-auth/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetProfile(ctx *gin.Context) {
	uid, _ := utils.ExtractTokenUID(ctx)

	user, err := models.GetProfile(uid)

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "user does not exists"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success", "data": user})

}
