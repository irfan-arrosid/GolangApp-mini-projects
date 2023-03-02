package models

type Product struct {
	Id   int64  `gorm:"primaryKey" json:"id"`
	Name string `gorm:"type:varchar(255)" json:"productName"`
	desc string `gorm:"type:text" json:"desc"`
}
