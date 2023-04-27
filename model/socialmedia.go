package model

import (
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type SocialMedia struct {
	gorm.Model
	Name           string `gorm:"not null" json:"name" validate:"required"`
	SocialMediaUrl string `gorm:"not null" json:"social_media_url" validate:"required"`
	UserID         uint
	User           *User
}

func (sm *SocialMedia) BeforeCreate(tx *gorm.DB) error {
	validate := validator.New()

	err := validate.Struct(sm)
	return err
}
