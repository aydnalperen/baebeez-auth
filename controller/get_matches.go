package controller

import (
	"baebeez-auth/models"
	"baebeez-auth/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetMatches(ctx *gin.Context) {
	uid, _ := utils.ExtractTokenUID(ctx)
	var matches models.Match

	if result := models.DB.Model(&models.Match{}).Where("uid=?", uid).Find(&matches); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}
	ctx.JSON(http.StatusOK, &matches)
}
