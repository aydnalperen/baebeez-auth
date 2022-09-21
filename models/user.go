package models

import (
	"baebeez-auth/utils"
	"errors"

	"github.com/gin-gonic/gin"
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
	MatchInfo  string `gorm:"size:255;not null;" json:"match_info"`
	IsApproved int    `gorm:"default:0"`
}
type UserAuth struct {
	gorm.Model
	Uid             string `gorm:"size:255;not null;unique" json:"uid"`
	Mail            string `gorm:"size:255;not null;unique" json:"mail" binding:"required"`
	Password        string `json:"password" binding:"required"`
	IsVerified      int    `gorm:"default:0;"`
	LastConnectedIp string `json:"ip"`
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

func LoginCheck(mail string, password string, c *gin.Context) (string, error) {
	var err error

	user := UserAuth{}

	err = DB.Model(UserAuth{}).Where("e_mail=?", mail).Take(&user).Error

	if err != nil {
		return "", err
	}
	err = VerifyPassword(password, user.Password)

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}
	token, err := utils.GenerateToken(user.Uid, c)

	if err != nil {
		return "", err
	}

	return token, nil
}
func GetUserByUid(Uid string) (UserAuth, error) {
	var user UserAuth

	if err := DB.Model(UserAuth{}).Where("uid=?", Uid).Take(&user).Error; err != nil {
		return user, errors.New("User not found!")
	}

	return user, nil
}
