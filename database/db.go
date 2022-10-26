package database

import (
	"log"

	"github.com/go-sql-api/core/model"
	"github.com/go-sql-api/env"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DBConnect   *gorm.DB
	DB_USERNAME = env.ENV("DB_USERNAME")
	DB_PASSWORD = env.ENV("DB_PASSWORD")
	DB_PORT     = env.ENV("DB_HOST")
	DB_NAME     = env.ENV("DB_NAME")
)

func EstablishedDBConnect() {
	// connString := fmt.Sprintf(
	// 	"port=%s port=%s user=%s password=%s dbname=%s ",
	// 	DB_PORT, DB_USERNAME, DB_PASSWORD, DB_NAME,
	// )
	dns := DB_USERNAME + ":" + DB_PASSWORD + "@tcp(" + DB_PORT + ")" + "/" + DB_NAME + "?charset=utf8mb4&parseTime=True&loc=Local"
	conn, err := gorm.Open(mysql.Open(dns), &gorm.Config{})

	if err != nil {

		panic("Error")
	}
	log.Println("DB connected successfully")
	conn.AutoMigrate(&model.Product{})
	conn.AutoMigrate(&model.Customer{})
	conn.AutoMigrate(&model.Company{})
	DBConnect = conn

}
