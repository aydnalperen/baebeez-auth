package controller

import (
	"baebeez-auth/models"
	"baebeez-auth/utils"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func GetUsers(ctx *gin.Context) {
	uid, _ := utils.ExtractTokenUID(ctx)
	var users []models.User

	if result := models.DB.Model(&models.User{}).Joins("inner join likes on likes.source=? and users.uid!=likes.destination and users.uid != likes.source", uid).Scan(&users); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return

	}

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(users), func(i, j int) { users[i], users[j] = users[j], users[i] })

	var return_users []models.User
	for i := 0; i < len(users); i++ {
		return_users = append(return_users, users[i])
	}
	ctx.JSON(http.StatusOK, &return_users)
}
