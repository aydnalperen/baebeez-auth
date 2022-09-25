package models

type ProfileInput struct {
	Name  string `gorm:"size:255;not null;" json:"name"`
	Photo string `gorm:"size:255;unique" json:"photo"`
	Major string `gorm:"size:255;not null;" json:"major"`
	Year  int    `gorm:"not null;" json:"year"`
	Bio   string `gorm:"not null;" json:"bio"`
	Uid   string `gorm:"size:255;not null;unique" json:"uid"`
}
