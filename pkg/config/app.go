package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	database *gorm.DB
)

func Main() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root@tcp(localhost:3306)/goBookstore?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	database = db
}

func GetDB() *gorm.DB {
	return database
}
