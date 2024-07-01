package config

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() {
	d, err := gorm.Open("mysql", "user:password@tcp(127.0.0.1:3306)/mydatabase?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
