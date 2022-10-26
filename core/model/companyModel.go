package model

import "gorm.io/gorm"

type Company struct {
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Adress   string `json:"adress"`
	Personal int    `json:"personal"`
	gorm.Model
}
