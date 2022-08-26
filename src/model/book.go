package model

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Name        string   `json:"name" gorm:"type: varchar(200)" validate:"required"`
	Description string   `json:"description" gorm:"type: varchar(200)" validate:"required"`
	Author      string   `json:"author" gorm:"type: varchar(100)" validate:"required"`
	Rate        int      `json:"rate"`
	Price       int      `json:"price" validate:"required"`
	CategoryId  int      `json:"categoryid" validate:"required"`
	Picture     string   `json:"picture" gorm:"type: varchar(4000)" `
	Catogory    Category `gorm:"references:categoryid"`
}
