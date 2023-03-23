package models

import (
	"gorm.io/gorm"
)

//Foreign key is USER ID
type Post struct {
	gorm.Model
	ID          int64  `gorm:"primarykey"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Comments    []Comment
	Likes       []Like
	UserID      int64
	User        User
}

func (post *Post) CreatePost(id int64) (*Post, *gorm.DB) {
	post.UserID = id
	if db := db.Create(&post); db.Error != nil {
		return nil, db
	}
	post, db := GetPostById(post.ID)

	return post, db
}

func GetAllPosts() []Post {
	var posts []Post
	db.Model(Post{}).Preload("User").Preload("Comments").Preload("Likes").Find(&posts)
	return posts
}

func GetPostById(Id int64) (*Post, *gorm.DB) {
	var post Post
	if db := db.Preload("User").Preload("Comments").Preload("Likes").Where("ID=?", Id).Find(&post).First(&post); db.Error != nil {
		return nil, db
	}
	return &post, db
}

func UpdatePost(Id int64, updatedPost *Post) (*Post, *gorm.DB) {
	if db := db.Model(Post{}).Where("ID = ?", Id).Updates(&updatedPost).First(&updatedPost); db.Error != nil {
		return nil, db
	}
	post, db := GetPostById(Id)
	return post, db
}

func DeletePost(Id int64) (*Post, *gorm.DB) {
	var post Post
	if db := db.Preload("User").Where("ID=?", Id).Find(&post).First(&post).Delete(&post); db.Error != nil {
		return nil, db
	}
	db.Preload("User").Unscoped().Where("ID=?", Id).Find(&post)
	return &post, db
}
