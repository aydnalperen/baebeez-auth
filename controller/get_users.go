package controller

import (
	"baebeez-auth/models"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func GetUsers(ctx *gin.Context) {
	// uid, _ := utils.ExtractTokenUID(ctx)
	var users []models.User

	// query := "SELECT * FROM users INNER JOIN matches ON users.uid = matches.person1 "
	if result := models.DB.Find(&users); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(users), func(i, j int) { users[i], users[j] = users[j], users[i] })

	var return_users []models.User
	for i := 0; i < 10; i++ {
		return_users = append(return_users, users[i])
	}
	ctx.JSON(http.StatusOK, &return_users)
}
