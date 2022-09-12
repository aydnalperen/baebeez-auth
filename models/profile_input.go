package models

type ProfileInput struct {
	FirstName  string `gorm:"size:255;not null;" json:"firstname"`
	LastName   string `gorm:"size:255;not null;" json:"lastname"`
	Photo      string `gorm:"size:255;not null;unique" json:"photo"`
	Major      string `gorm:"size:255;not null;" json:"major"`
	Year       int    `gorm:"not null;" json:"year"`
	Bio        string `gorm:"not null;" json:"bio"`
	Department string `gorm:"not null;" json:"department"`
	Password   string `gorm:"size:255;not null;unique" json:"password"`
	Email      string `gorm:"size:255;not null;unique" json:"email"`
}
