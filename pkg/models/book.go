package models

import (
	"gorm.io/gorm"
)

//Foreign key is USER ID
type Book struct {
	gorm.Model
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
	UserID      uint
}

func (book *Book) CreateBook(id uint) (*Book, *gorm.DB) {
	book.UserID = id
	if db := db.Create(&book); db.Error != nil {
		return nil, db
	}
	return book, db
}

func GetAllBooks() []Book {
	var Books []Book
	db.Model(Book{}).Find(&Books)
	return Books
}

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var book Book
	if db := db.Where("ID=?", Id).Find(&book).First(&book); db.Error != nil {
		return nil, db
	}
	return &book, db
}

func UpdateBook(Id int64, updatedBook *Book) (*Book, *gorm.DB) {
	if db := db.Model(Book{}).Where("ID = ?", Id).Updates(&updatedBook).First(&updatedBook); db.Error != nil {
		return nil, db
	}
	b, db := GetBookById(Id)
	return b, db
}

func DeleteBook(Id int64) (*Book, *gorm.DB) {
	var book Book
	if db := db.Where("ID=?", Id).Find(&book).First(&book).Delete(&book); db.Error != nil {
		return nil, db
	}
	db.Unscoped().Where("ID=?", Id).Find(&book)
	return &book, db
}
