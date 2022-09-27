package models

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

var EmailDomainCheck validator.Func = func(fl validator.FieldLevel) bool {
	emailhost := "@itu.edu.tr"
	return strings.Contains(fl.Field().String(), emailhost)

}

type LoginInput struct {
	Mail     string `json:"mail" binding:"required,email,EmailDomainCheck"`
	Password string `json:"password" binding:"required"`
}
