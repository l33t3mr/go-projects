package model

import "gorm.io/gorm"

type Director struct {
	gorm.Model
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	ClipId    uint
	Clips     []*Clip `gorm:"many2many:director_clips;" json:"clips,omitempty"`
}

func (d Director) Migrate(db *gorm.DB) {
	db.AutoMigrate(&d)
}
