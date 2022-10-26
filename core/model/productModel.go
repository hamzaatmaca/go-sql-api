package model

import "gorm.io/gorm"

type Product struct {
	Name  string `json:"name"`
	Brand string `json:"brand"`
	Price int    `json:"price"`
	gorm.Model
}
