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

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var book Book
	db := db.Where("ID=?", Id).Find(&book)
	return &book, db
}

func UpdateBook(Id int64, updatedBook *Book) (*Book, *gorm.DB) {
	db := db.Model(Book{}).Where("ID = ?", Id).Updates(&updatedBook)
	b, _ := GetBookById(Id)
	return b, db
}

func DeleteBook(ID int64) (*Book, *gorm.DB) {
	var book Book
	db.Where("ID=?", ID).Delete(&book)
	db := db.Unscoped().Where("ID=?", ID).Find(&book)
	return &book, db
}
