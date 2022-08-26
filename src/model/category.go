package model

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name        string `json:"name" gorm: "type:varchar(200)" validate:"required"`
	Description string `json:"description" gorm: "type:varchar(2000)" validate:"required"`
}
