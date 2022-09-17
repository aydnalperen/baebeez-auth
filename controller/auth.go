package controller

import (
	"baebeez-auth/models"
	"baebeez-auth/utils"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func Register(ctx *gin.Context) {
	var input models.RegisterInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.UserAuth

	user.EMail = input.EMail
	user.Uid = uuid.NewString()
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)

	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
	}
	user.Password = string(hashedPassword) // user creation is done

	_, err = user.SaveUserAuth()

	verifCode := new(models.VerifCode)

	verifCode.Uid = user.Uid
	verifCode.VerifCode = models.GetRandomFourDigit()

	verifCode.SaveVerifCode()

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "registered!", "mail": user.EMail, "password": user.Password})
}
func SaveProfile(ctx *gin.Context) {
	var input models.ProfileInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User

	user.Bio = input.Bio
	user.Department = input.Department
	user.FirstName = input.FirstName
	user.LastName = input.LastName
	user.Major = input.Major
	user.Photo = input.Photo
	user.Year = input.Year
	user.Uid = uuid.NewString()

	_, err := user.SaveUser()

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

	user, err := models.GetUserByUid(user_id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success", "data": user})
}
