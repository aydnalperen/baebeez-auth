package models

import (
	"math/rand"
	"strconv"

	"github.com/jinzhu/gorm"
)

type VerifCode struct {
	gorm.Model
	Uid              string `gorm:"size:255;not null;unique" json:"uid"`
	VerifCode        string `gorm:"size:255;not null;unique" json:"verif_code"`
	CodeEntryCounter int    `gorm:"default:0" json:"entry_counter"`
}

func GetRandomFourDigit() string {
	min := 1000
	max := 9999
	var r = int64(rand.Intn(max-min) + min)
	return strconv.FormatInt(r, 10)
}
