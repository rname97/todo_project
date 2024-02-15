package config

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	// bisa disesuaikan dengan config database sendiri
	username := "root"
	password := "root"
	dbName := "todolist"
	dbHost := "localhost"
	dbPort := "8889"

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", username, password, dbHost, dbPort, dbName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}

	DB = db

	// migrasi model
	// db.AutoMigrate(&Task{})
}
