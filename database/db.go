package database

import (
	"fmt"
	"log"

	"github.com/suryasatriah/mygram/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	dbHost     = "localhost"
	dbUser     = "postgres"
	dbPassword = "12345"
	dbName     = "mygram"
	dbPort     = "5432"
)

var (
	db  *gorm.DB
	err error
)

func ConnectDatabase() {

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", dbHost, dbUser, dbPassword, dbName, dbPort)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("Cannot connect to database, ", err)
	}
	log.Println("Connected to database.")

	db.Debug().AutoMigrate(model.User{}, model.SocialMedia{}, model.Photo{}, model.Comment{})

}

func GetDatabase() *gorm.DB {
	return db
}
