package models

import "github.com/jinzhu/gorm"

type Match struct {
	gorm.Model
	Person1 string `gorm:"size:255;not null;"`
	Person2 string `gorm:"size:255;not null;"`
}

func NewMatch(src string, dest string) *Match {
	l := new(Match)
	l.Person1 = src
	l.Person2 = dest

	return l
}
