package validations

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var EmailDomainCheck validator.Func = func(fl validator.FieldLevel) bool {
	return strings.HasSuffix(fl.Field().String(), "@itu.edu.tr")

}

type LoginSchema struct {
	Mail     string `json:"mail" binding:"required,email,EmailDomainCheck"`
	Password string `json:"password" binding:"required,min=8,max=20"`
}

func LoginValidation(c *gin.Context) {
	var input LoginSchema
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}
	c.Set("loginInput", input)
	c.Next()
}
