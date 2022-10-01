package controller

import (
	"baebeez-auth/models"
	"baebeez-auth/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddToLikes(ctx *gin.Context) {
	destUid := ctx.Param("dest_uid")

	source, err := utils.ExtractTokenUID(ctx)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
	}
	like := models.NewLike(source, destUid)

	models.DB.Save(like)
	//after create ile yapılabilir her like eklendikten sonra 
	//otomatik olarak match var mı diye kontrol ettirebiliriz. GORM-HOOKS
	if CheckMatch(source, destUid) {
		match := models.NewMatch(source, destUid)

		models.DB.Save(match)
		ctx.JSON(http.StatusOK, gin.H{"match": 1})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"match": 0})
	}
}
func CheckMatch(source string, destination string) bool {
	like2 := new(models.Like)

	models.DB.Model(&models.Like{}).Where("source = ? AND destination = ?", destination, source).Find(&like2)

	if like2 != nil {
		return true
	}

	return false
}
