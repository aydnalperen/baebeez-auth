package models

import "github.com/jinzhu/gorm"

type Match struct {
	gorm.Model
	person1 string `gorm:"not null;"`
	person2 string `gorm:"not null;"`
}

func NewMatch(src string, dest string) *Match {
	l := new(Match)
	l.person1 = src
	l.person2 = dest

	return l
}
