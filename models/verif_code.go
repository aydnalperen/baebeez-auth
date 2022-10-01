package models

import (
	"crypto/rand"
	"io"

	"github.com/jinzhu/gorm"
)

type VerifCode struct {
	gorm.Model
	Uid              string `gorm:"size:255;not null;unique" json:"uid"`
	SendDate         string `gorm:"size:255" json:"send_date"`
	VerifCode        string `gorm:"size:255;not null;" json:"verif_code"`
	CodeEntryCounter int    `gorm:"default:0" json:"entry_counter"`
}

func (v *VerifCode) SaveVerifCode() (*VerifCode, error) {
	if result := DB.Create(&v); result.Error != nil {
		return &VerifCode{}, result.Error
	}
	return v, nil
}

func GetLastVerifCodeByUid(uid string) string {
	var verifcode VerifCode

	if result := DB.Model(&VerifCode{}).Where("uid=?", uid).Order("date DESC").First(&verifcode); result != nil {
		return " "
	}

	return verifcode.VerifCode
}
func CreateVerifCode(max int) string {
	b := make([]byte, max)
	n, err := io.ReadAtLeast(rand.Reader, b, max)
	if n != max {
		panic(err)
	}
	for i := 0; i < len(b); i++ {
		b[i] = table[int(b[i])%len(table)]
	}
	return string(b)
}

var table = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}
