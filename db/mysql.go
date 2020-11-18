package db

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() (err error) {
	USER := "root"
	PASS := "root123"
	HOST := "127.0.0.1"
	PORT := "3306"
	DBNAME := "todoDB"

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		USER,
		PASS,
		HOST,
		PORT,
		DBNAME,
	)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	return nil
}
