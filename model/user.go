package model

import (
	"github.com/go-playground/validator/v10"
	"github.com/suryasatriah/mygram/helper"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `gorm:"not null,uniqueIndex" json:"email" validate:"required,email"`
	Username string `gorm:"not null,unique" json:"username" validate:"required"`
	Password string `gorm:"not null" json:"password" validate:"required"`
	Age      int    `gorm:"not null" json:"age" validate:"required"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	validate := validator.New()

	err := validate.Struct(u)
	if err != nil {
		return err
	}
	
	u.Password = helper.HashPassword(u.Password)
	return err
}