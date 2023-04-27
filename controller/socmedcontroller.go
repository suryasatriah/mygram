package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/suryasatriah/mygram/database"
	"github.com/suryasatriah/mygram/helper"
	"github.com/suryasatriah/mygram/model"
)

func GetSocialMedia(c *gin.Context) {
	db := database.GetDatabase()
	contentType := helper.GetContentType(c)
	SocialMedia := model.SocialMedia{}

	if contentType == appJSON {
		c.ShouldBindJSON(&SocialMedia)
	} else {
		c.ShouldBind(&SocialMedia)
	}

	err := db.Debug().Find(&SocialMedia).Error

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, SocialMedia)
}

func CreateSocialMedia(c *gin.Context) {
	db := database.GetDatabase()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helper.GetContentType(c)

	SocialMedia := model.SocialMedia{}
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&SocialMedia)
	} else {
		c.ShouldBind(&SocialMedia)
	}

	SocialMedia.UserID = userID

	err := db.Debug().Create(&SocialMedia).Error
	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, SocialMedia)
}
