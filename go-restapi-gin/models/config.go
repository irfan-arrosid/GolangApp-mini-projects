package models

import (
	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
)

// var DB *gorm.DB

// func dbConnect() {
// 	dsn := "postgres://irfanarrosid:at19ir97ar@localhost:5432/restapi-gin"
// 	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

// 	if err != nil {
// 		panic("Failed to connect database")
// 	}

// 	db.AutoMigrate(&Product{})

// 	DB = db
// }

type Product struct {
	Id   int64  `gorm:"primaryKey" json:"id"`
	Name string `gorm:"type:varchar(255)" json:"productName"`
	desc string `gorm:"type:text" json:"desc"`
}

func dbConnect() {
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=irfanarrosid dbname=restapi-gin password=at19ir97ar sslmode=disable")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	var Products []Product
	db.Find(&Products)
}
