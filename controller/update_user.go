package controller

import (
	"baebeez-auth/models"
	"baebeez-auth/utils"
	"baebeez-auth/validations"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UpdateUserRequestBody struct {
	Name  string `gorm:"size:255;not null;" json:"name"`
	Photo string `gorm:"size:255;not null;unique" json:"photo"`
	Major string `gorm:"size:255;not null;" json:"major"`
	Year  int    `gorm:"not null;" json:"year"`
	Bio   string `gorm:"not null;" json:"bio"`
}

func UpdateUser(ctx *gin.Context) {
	uid, _ := utils.ExtractTokenUID(ctx)

	// body := UpdateUserRequestBody{}
	var user models.User

	// if err := ctx.BindJSON(&body); err != nil {
	// 	ctx.AbortWithError(http.StatusBadRequest, err)
	// 	return
	// }
	data, _ := ctx.Get("UpdateUserInput")

	body := data.(validations.UpdateUserSchema)

	if result := models.DB.First(&user, uid); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	user.Bio = body.Bio

	user.Year = body.Year
	user.Photo = body.Photo
	user.Name = body.Name

	user.Major = body.Major

	models.DB.Save(&user)

	ctx.JSON(http.StatusOK, &user)

}
