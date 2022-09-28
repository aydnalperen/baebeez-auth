package models

type RegisterInput struct {
	LoginInput
	// Mail     string `json:"mail" binding:"required"`
	// Password string `json:"password" binding:"required"`
}
