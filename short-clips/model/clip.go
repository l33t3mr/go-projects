package model

import (
	"gorm.io/gorm"
)

type Clip struct {
	gorm.Model
	ISBN      string      `json:"isbn"`
	Title     string      `json:"title"`
	Directors []*Director `gorm:"many2many:director_clips;" json:"directors,omitempty"`
}

func (c Clip) Migrate(db *gorm.DB) {
	db.AutoMigrate(&c)
}
