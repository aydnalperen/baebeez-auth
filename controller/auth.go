package controller

import (
	"go-authapi-adv/models"
	"go-authapi-adv/utils"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"github.com/gin-gonic/gin"
)

func Register(ctx *gin.Context) {
	var input models.RegisterInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.UnVerifiedUser

	user.EMail = input.EMail

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)

	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
	}
	user.Password = string(hashedPassword) // user creation is done

	_, err = user.SaveUnVerifiedUser()

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "validated!"})
}

func Login(ctx *gin.Context) {
	var input models.LoginInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "username or password is incorrect."})
	}
	token, err := models.LoginCheck(input.Mail, input.Password)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "username or password is incorrect."})
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token})
}

func Logout(ctx *gin.Context) {

}
func CurrentUser(ctx *gin.Context) {
	user_id, err := utils.ExtractTokenID(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	user, err := models.GetUserById(user_id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success", "data": user})
}
