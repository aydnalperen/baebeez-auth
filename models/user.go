package models

import (
	"errors"
	"go-authapi-adv/utils"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model
	Uid        string `gorm:"size:255;not null;unique" json:"uid"`
	Email      string `gorm:"size:255;not null;unique" json:"email"`
	FirstName  string `gorm:"size:255;not null;" json:"firstname"`
	LastName   string `gorm:"size:255;not null;" json:"lastname"`
	Password   string `gorm:"size:255;not null;unique" json:"password"`
	Photo      string `gorm:"size:255;not null;unique" json:"photo"`
	Major      string `gorm:"size:255;not null;" json:"major"`
	Year       int    `gorm:"not null;" json:"year"`
	Bio        string `gorm:"not null;" json:"bio"`
	Department string `gorm:"not null;" json:"department"`
}
type UnVerifiedUser struct {
	EMail    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
func (u *User) SaveUser() (*User, error) {
	if result := DB.Create(&u); result.Error != nil {
		return &User{}, result.Error
	}
	return u, nil
}
func (u *UnVerifiedUser) SaveUnVerifiedUser() (*UnVerifiedUser, error) {
	if result := DB.Create(&u); result.Error != nil {
		return &UnVerifiedUser{}, result.Error
	}
	return u, nil
}

func LoginCheck(username string, password string) (string, error) {
	var err error

	user := User{}

	err = DB.Model(User{}).Where("username=?", username).Take(&user).Error

	if err != nil {
		return "", err
	}
	err = VerifyPassword(password, user.Password)

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}
	token, err := utils.GenerateToken(user.ID)

	if err != nil {
		return "", err
	}

	return token, nil
}
func GetUserById(id uint) (User, error) {
	var user User

	if err := DB.First(&user, id).Error; err != nil {
		return user, errors.New("User not found!")
	}

	user.Password = ""

	return user, nil
}
