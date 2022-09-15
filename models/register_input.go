package models

type RegisterInput struct {
	EMail    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
