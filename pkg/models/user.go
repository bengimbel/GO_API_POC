package models

import (
	"log"

	"github.com/bengimbel/go-bookstore/pkg/config"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var database *gorm.DB

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Username string `json:"username" gorm:"unique"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
}

func init() {
	config.Main()
	database = config.GetDB()
	database.AutoMigrate(&User{})
}

func (user *User) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	log.Println(user, "HAS")
	return nil
}
func (user *User) CheckPassword(providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(providedPassword))
	if err != nil {
		return err
	}
	return nil
}

func (user *User) CreateUser(newUser *User) *User {
	db.Create(&user)
	return user
}

func (user *User) GetUserByEmail(email string) (*User, *gorm.DB, error) {
	if db := db.Where("email=?", email).Find(&user).First(&user); db.Error != nil {
		return nil, db, db.Error
	}
	return user, db, nil
}
