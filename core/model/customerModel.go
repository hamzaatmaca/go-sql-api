package model

import "gorm.io/gorm"

type Customer struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Age     int    `json:"age"`
	Salary  int    `json:"salary"`
	gorm.Model
}
