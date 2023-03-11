package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connection() *gorm.DB {
	dsn := "root:1234@tcp(127.0.0.1:3306)/wallet?parseTime=True"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Connection Failed")
	}
	return db
}
