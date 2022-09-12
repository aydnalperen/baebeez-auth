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
	Photo      string `gorm:"size:255;not null;unique" json:"photo"`
	Major      string `gorm:"size:255;not null;" json:"major"`
	Year       int    `gorm:"not null;" json:"year"`
	Bio        string `gorm:"not null;" json:"bio"`
	Department string `gorm:"not null;" json:"department"`
	FirstName  string `gorm:"size:255;not null;" json:"firstname"`
	LastName   string `gorm:"size:255;not null;" json:"lastname"`
	IsComp     int    `gorm:"default:0"`
}
type UserAuth struct {
	gorm.Model
	Uid        string `gorm:"size:255;not null;unique" json:"uid"`
	EMail      string `json:"email" binding:"required"`
	Password   string `json:"password" binding:"required"`
	IsVerified int    `gorm:"default:0;"`
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
func (u *UserAuth) SaveUserAuth() (*UserAuth, error) {
	if result := DB.Create(&u); result.Error != nil {
		return &UserAuth{}, result.Error
	}
	return u, nil
}

func LoginCheck(username string, password string) (string, error) {
	var err error

	user := UserAuth{}

	err = DB.Model(User{}).Where("username=?", username).Take(&user).Error

	if err != nil {
		return "", err
	}
	err = VerifyPassword(password, user.Password)

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}
	token, err := utils.GenerateToken(user.Uid)

	if err != nil {
		return "", err
	}

	return token, nil
}
func GetUserByUid(Uid string) (User, error) {
	var user User

	if err := DB.First(&user, Uid).Error; err != nil {
		return user, errors.New("User not found!")
	}

	return user, nil
}
