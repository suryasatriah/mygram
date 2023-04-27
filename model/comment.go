package model

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	UserID  uint
	PhotoID uint
	Message string `gorm:"not null" json:"message" validate:"required"`
	User    User
	Photo   Photo
}