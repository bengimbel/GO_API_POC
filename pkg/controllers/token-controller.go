package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/bengimbel/go-bookstore/pkg/auth"
	"github.com/bengimbel/go-bookstore/pkg/models"
	"github.com/bengimbel/go-bookstore/pkg/utils"
)

type TokenRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func GenerateToken(w http.ResponseWriter, r *http.Request) {
	User := &models.User{}

	RequestObj := &TokenRequest{}
	utils.ParseBody(r, RequestObj)
	// if err := context.ShouldBindJSON(&request); err != nil {
	// 	context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	context.Abort()
	// 	return
	// }
	// check if email exists and password is correct
	_, _, err := User.GetUserByEmail(RequestObj.Email)

	if err != nil {
		log.Println("CANNOT FIND EMAIL")
		// context.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		// context.Abort()
		return
	}
	credentialError := User.CheckPassword(RequestObj.Password)
	if credentialError != nil {
		log.Println("BAD PASSWORD")
		// context.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		// context.Abort()
		return
	}
	tokenString, err := auth.GenerateJWT(User.Email, User.Username)
	if err != nil {
		// context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		// context.Abort()
		log.Println("CANNOT MAKE TOKEN")
		return
	}
	// context.JSON(http.StatusOK, gin.H{"token": tokenString})
	returnV := map[string]string{"token": tokenString}
	res, _ := json.Marshal(returnV)
	// b := CreateBook.CreateBook()
	// res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
