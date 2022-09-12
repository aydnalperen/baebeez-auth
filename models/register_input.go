package models

type RegisterInput struct {
	EMail    string `json:"mail" binding:"required"`
	Password string `json:"password" binding:"required"`
}
