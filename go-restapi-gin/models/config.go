package models

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func dbConnect() {
	dsn := "postgres://irfanarrosid:at19ir97ar@localhost:5432/restapi-gin"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect database")
	}

	db.AutoMigrate(&Product{})

	DB = db
}
