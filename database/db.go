package database

import (
	"api-go-gin/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
	err error
)

func ConnectDB()  {
	connection := "host=localhost user=postgres password=root dbname=postgres port=5432 sslmode=disable"
	con, err := gorm.Open(postgres.Open(connection))
	if err != nil {
		log.Panic(err.Error())
	}

	con.AutoMigrate(&models.Student{})
	DB = con
}
