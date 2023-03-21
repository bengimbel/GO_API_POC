package models

import (
	"github.com/bengimbel/go-bookstore/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	config.Main()
	db = config.GetDB()
	db.Migrator().CreateConstraint(&Book{}, "Users")
	db.AutoMigrate(&User{}, &Book{})
}
