package models

import (
	"github.com/bengimbel/go-bookstore/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Main()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	db.Create(&b)
	return b
}

func GetAllBooks() []Book {
	var Books []Book
	db.Model(Book{}).Find(&Books)
	return Books
}

func GetBookById(Id int64) (*Book, *gorm.DB, error) {
	var book Book
	if db := db.Where("ID=?", Id).Find(&book).First(&book); db.Error != nil {
		return nil, db, db.Error
	}
	return &book, db, nil
}

func UpdateBook(Id int64, updatedBook *Book) (*Book, *gorm.DB, error) {
	if db := db.Model(Book{}).Where("ID = ?", Id).Updates(&updatedBook).First(&updatedBook); db.Error != nil {
		return nil, db, db.Error
	}
	b, _, _ := GetBookById(Id)
	return b, db, nil
}

func DeleteBook(Id int64) (*Book, *gorm.DB, error) {
	var book Book
	if db := db.Where("ID=?", Id).Find(&book).First(&book).Delete(&book); db.Error != nil {
		return nil, db, db.Error
	}
	db.Unscoped().Where("ID=?", Id).Find(&book)
	return &book, db, nil
}
