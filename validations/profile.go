package validations

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var MajorCheck validator.Func = func(fl validator.FieldLevel) bool {
	return true
}

type ProfileSchema struct {
	Name  string `json:"name" binding:"required"`
	Photo string `json:"photo" binding:"required"`
	Major string `json:"major" binding:"required"`
	Year  int    `json:"year" binding:"required"`
	Bio   string `json:"bio" binding:"required"`
	Uid   string `json:"uid" binding:"required"`
}

func ProfileValidation(c *gin.Context) {
	var input ProfileSchema
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}
	c.Set("ProfileInput", input)
	c.Next()
}
