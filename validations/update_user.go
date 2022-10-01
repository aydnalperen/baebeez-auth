package validations

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UpdateUserSchema struct {
	Name  string `json:"name" binding:"required"`
	Photo string `json:"photo" binding:"required"`
	Major string `json:"major" binding:"required"`
	Year  int    `json:"year" binding:"required"`
	Bio   string `json:"bio" binding:"required"`
}

func UpdateUserValidation(c *gin.Context) {
	var input UpdateUserSchema
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}
	c.Set("UpdateUserInput", input)
	c.Next()
}
