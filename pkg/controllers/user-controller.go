package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/bengimbel/go-bookstore/pkg/models"
	"github.com/bengimbel/go-bookstore/pkg/utils"
)

var User models.User

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	NewUser := &models.User{}
	utils.ParseBody(r, NewUser)
	if err := NewUser.HashPassword(NewUser.Password); err != nil {
		errorMap := map[string]string{"error": "Error creating user", "code": "500"}
		errorJson, _ := json.Marshal(errorMap)
		w.Header().Set("Content-Type", "pkglication/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(errorJson)
		return
	}
	CreatedUser, db := NewUser.CreateUser()
	if db.Error != nil {
		errorMap := map[string]string{"error": "Error creating user", "code": "500"}
		errorJson, _ := json.Marshal(errorMap)
		w.Header().Set("Content-Type", "pkglication/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(errorJson)
		return
	}
	returnV := map[string]string{"username": CreatedUser.Username, "email": CreatedUser.Email}
	res, _ := json.Marshal(returnV)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
