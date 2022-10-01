package models

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

var EmailDomainCheck validator.Func = func(fl validator.FieldLevel) bool {
	return strings.Contains(fl.Field().String(), "@itu.edu.tr")

}

type LoginInput struct {
	Mail     string `json:"mail" binding:"required"`
	Password string `json:"password" binding:"required,min=8,max=20"`
}
