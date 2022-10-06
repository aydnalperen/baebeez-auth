package validations

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type RegisterSchema struct {
	LoginSchema
}

func RegisterValidation(c *gin.Context) {
	var input RegisterSchema
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}
	c.Set("RegisterInput", input)
	c.Next()
}
