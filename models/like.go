package models

import "github.com/jinzhu/gorm"

type Like struct {
	gorm.Model
	Source      string `gorm:"size:255;not null;" json:"source"`
	Destination string `gorm:"size:255;not null;" json:"destination"`
}

func NewLike(src string, dest string) *Like {
	l := new(Like)
	l.Source = src
	l.Destination = dest

	return l
}
