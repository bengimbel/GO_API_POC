package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/bengimbel/go-bookstore/pkg/models"
	"github.com/bengimbel/go-bookstore/pkg/utils"
)

var User models.User

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	NewUser := &models.User{}
	utils.ParseBody(r, NewUser)
	log.Println(NewUser, "new user")
	log.Println(r, "new rrr")
	if err := NewUser.HashPassword(NewUser.Password); err != nil {
		errorMap := map[string]string{"error": "HAS FAIL", "code": "404"}
		errorJson, _ := json.Marshal(errorMap)
		w.Header().Set("Content-Type", "pkglication/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(errorJson)
		return
	}
	user := NewUser.CreateUser(NewUser)
	log.Println(user, "USA")
	returnV := map[string]string{"username": NewUser.Username, "email": NewUser.Email}
	res, _ := json.Marshal(returnV)
	// b := CreateBook.CreateBook()
	// res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

	// if err := context.ShouldBindJSON(&user); err != nil {
	// 	context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	context.Abort()
	// 	return
	// }
	// if err := user.HashPassword(user.Password); err != nil {
	// 	context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 	context.Abort()
	// 	return
	// }
	// record := database.Instance.Create(&user)
	// if record.Error != nil {
	// 	context.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
	// 	context.Abort()
	// 	return
	// }
	// context.JSON(http.StatusCreated, gin.H{"userId": user.ID, "email": user.Email, "username": user.Username})
}
