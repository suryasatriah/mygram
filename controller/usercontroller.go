package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/suryasatriah/mygram/database"
	"github.com/suryasatriah/mygram/helper"
	"github.com/suryasatriah/mygram/model"
)

var (
	appJSON = "application/json"
)

func UserRegister(c *gin.Context) {
	db := database.GetDatabase()
	contentType := helper.GetContentType(c)
	User := model.User{}

	if contentType == appJSON {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	err := db.Debug().Create(&User).Error
	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":       User.ID,
		"emai":     User.Email,
		"username": User.Username,
		"password": User.Password,
		"age":      User.Age,
	})

}

func UserLogin(c *gin.Context) {
	db := database.GetDatabase()
	contentType := helper.GetContentType(c)
	User := model.User{}
	password := ""

	if contentType == appJSON {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	password = User.Password

	err := db.Debug().Where("email = ?", User.Email).Take(&User).Error

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "invalid email/password",
		})
		return
	}

	comparePassword := helper.CompareHashedPassword([]byte(User.Password), []byte(password))

	if !comparePassword {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "invalid email/password",
		})
		return
	}

	token := helper.GenerateToken(User.ID, User.Email)

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})

}
