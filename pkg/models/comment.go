package models

import (
	"gorm.io/gorm"
)

//Foreign key is USER ID and POST ID
type Comment struct {
	gorm.Model
	ID      int64  `gorm:"primarykey"`
	Comment string `json:"comment"`
	PostID  int64
	UserID  int64
}

func (comment *Comment) CreateComment(id int64, postId int64) (*Comment, *gorm.DB) {
	comment.UserID = id
	comment.PostID = postId
	if db := db.Preload("Post").Create(&comment); db.Error != nil {
		return nil, db
	}
	return comment, db
}

func GetCommentById(Id int64) (*Post, *gorm.DB) {
	var post Post
	if db := db.Where("ID=?", Id).Find(&post).First(&post); db.Error != nil {
		return nil, db
	}
	return &post, db
}
