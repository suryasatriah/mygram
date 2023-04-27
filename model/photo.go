package model

import "gorm.io/gorm"

type Photo struct {
	gorm.Model
	Title    string `gorm:"not null" json:"title" validate:"required"`
	Caption  string `gorm:"not null" json:"caption" validate:"required"`
	PhotoUrl string `gorm:"not null" json:"photo_url" validate:"required,url"`
	UserID   uint
	User     User
}
