package controllers

import (
	"encoding/json"
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
	_, db := User.GetUserByEmail(RequestObj.Email)
	if db.Error != nil {
		errorMap := map[string]string{"error": "Email not found", "code": "404"}
		errorJson, _ := json.Marshal(errorMap)
		w.Header().Set("Content-Type", "pkglication/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(errorJson)
		return
	}
	credentialError := User.CheckPassword(RequestObj.Password)
	if credentialError != nil {
		errorMap := map[string]string{"error": "Incorrect password", "code": "401"}
		errorJson, _ := json.Marshal(errorMap)
		w.Header().Set("Content-Type", "pkglication/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(errorJson)
		return
	}
	tokenString, err := auth.GenerateJWT(User.Email, User.Username)
	if err != nil {
		errorMap := map[string]string{"error": "Error creating token", "code": "500"}
		errorJson, _ := json.Marshal(errorMap)
		w.Header().Set("Content-Type", "pkglication/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(errorJson)
		return
	}
	returnV := map[string]string{"token": tokenString}
	res, _ := json.Marshal(returnV)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
