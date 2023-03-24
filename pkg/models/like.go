package models

import (
	"gorm.io/gorm"
)

//Foreign key is USER ID and POST ID
type Like struct {
	gorm.Model
	PostID int64
	UserID int64
}

func (like *Like) CreateLike(id int64, postId int64) (*Like, *gorm.DB) {
	like.UserID = id
	like.PostID = postId
	if db := db.Preload("User").Preload("Post").Create(like); db.Error != nil {
		return nil, db
	}
	return like, db
}

func GetLikeById(Id int64) (*Like, *gorm.DB) {
	var like Like
	if db := db.Where("ID=?", Id).Find(&like).First(&like); db.Error != nil {
		return nil, db
	}
	return &like, db
}
