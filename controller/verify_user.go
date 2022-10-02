package controller

import (
	"baebeez-auth/models"
	"baebeez-auth/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func VerifyUser(ctx *gin.Context) {
	verif_code, is_exists := ctx.Params.Get("verif_code")

	if !is_exists {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "verification code required!"})
		return
	}

	uid, _ := utils.ExtractTokenUID(ctx)

	lastVerifCode := models.GetLastVerifCodeByUid(uid)

	models.IncreaseCounter(uid)
	if lastVerifCode == " " {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "verification code does not exist!"})
		return
	}

	if verif_code != lastVerifCode {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "verification code is invalid!"})
		return
	}

	if err := models.MakeUserValid(uid); err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.Status(http.StatusOK)

}
