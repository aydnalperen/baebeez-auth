package models

import "github.com/jinzhu/gorm"

type Like struct {
	gorm.Model
	source      string `gorm:"not null;"`
	destination string `gorm:"not null;"`
}

func NewLike(src string, dest string) *Like {
	l := new(Like)
	l.source = src
	l.destination = dest

	return l
}
