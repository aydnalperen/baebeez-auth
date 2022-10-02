package controller

import (
	"baebeez-auth/models"
	"baebeez-auth/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type VerifyUserRequestBody struct {
	VerifCode string `json:"verif_code"`
}

func VerifyUser(ctx *gin.Context) {

	var verifCode VerifyUserRequestBody
	if err := ctx.ShouldBindJSON(&verifCode); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "verification code required!"})
		return
	}

	verif_code := verifCode.VerifCode
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
