package main

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// var err error
	db, err := gorm.Open(postgres.Open(os.Getenv("DB_URL")), &gorm.Config{})

	if err != nil {
		fmt.Println("Failed to connect database!")
	} else {
		fmt.Println("Database connected....")
	}

	db.Raw("SELECT * FROM mens_shoes")
}
