package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/bengimbel/go-bookstore/pkg/models"
	"github.com/bengimbel/go-bookstore/pkg/utils"
	"github.com/gorilla/mux"
)

var NewPost models.Post

func GetPosts(w http.ResponseWriter, r *http.Request) {
	posts := models.GetAllPosts()
	res, _ := json.Marshal(posts)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetPostById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	postId := vars["postId"]
	Id, err := strconv.ParseInt(postId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	post, db := models.GetPostById(Id)
	if db.Error != nil {
		errorMap := map[string]string{"error": db.Error.Error(), "code": "404"}
		errorJson, _ := json.Marshal(errorMap)
		w.Header().Set("Content-Type", "pkglication/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(errorJson)
		return
	}
	res, _ := json.Marshal(post)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreatePost(w http.ResponseWriter, r *http.Request) {
	CreatePost := &models.Post{}
	utils.ParseBody(r, CreatePost)
	email := r.Header.Get("email")
	user, db := models.GetUserByEmail(email)
	post, db := CreatePost.CreatePost(user.ID)
	if db.Error != nil {
		errorMap := map[string]string{"error": db.Error.Error(), "code": "404"}
		errorJson, _ := json.Marshal(errorMap)
		w.Header().Set("Content-Type", "pkglication/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(errorJson)
		return
	}
	res, _ := json.Marshal(post)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateComment(w http.ResponseWriter, r *http.Request) {
	CreateComment := &models.Comment{}
	utils.ParseBody(r, CreateComment)
	vars := mux.Vars(r)
	id := vars["postId"]
	postId, err := strconv.ParseInt(id, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	post, db := models.GetPostById(postId)
	email := r.Header.Get("email")
	user, db := models.GetUserByEmail(email)
	comment, db := CreateComment.CreateComment(user.ID, post.ID)
	if db.Error != nil {
		errorMap := map[string]string{"error": db.Error.Error(), "code": "404"}
		errorJson, _ := json.Marshal(errorMap)
		w.Header().Set("Content-Type", "pkglication/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(errorJson)
		return
	}
	res, _ := json.Marshal(comment)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateLike(w http.ResponseWriter, r *http.Request) {
	CreateLike := &models.Like{}
	vars := mux.Vars(r)
	id := vars["postId"]
	postId, err := strconv.ParseInt(id, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	post, db := models.GetPostById(postId)
	email := r.Header.Get("email")
	user, db := models.GetUserByEmail(email)
	like, db := CreateLike.CreateLike(user.ID, post.ID)
	if db.Error != nil {
		errorMap := map[string]string{"error": db.Error.Error(), "code": "404"}
		errorJson, _ := json.Marshal(errorMap)
		w.Header().Set("Content-Type", "pkglication/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(errorJson)
		return
	}
	res, _ := json.Marshal(like)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdatePost(w http.ResponseWriter, r *http.Request) {
	UpdatePost := &models.Post{}
	utils.ParseBody(r, UpdatePost)
	vars := mux.Vars(r)
	postId := vars["postId"]
	Id, err := strconv.ParseInt(postId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	newPost, db := models.UpdatePost(Id, UpdatePost)
	if db.Error != nil {
		errorMap := map[string]string{"error": db.Error.Error(), "code": "404"}
		errorJson, _ := json.Marshal(errorMap)
		w.Header().Set("Content-Type", "pkglication/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(errorJson)
		return
	}
	res, _ := json.Marshal(newPost)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeletePost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	postId := vars["postId"]
	Id, err := strconv.ParseInt(postId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	deletedPost, db := models.DeletePost(Id)
	if db.Error != nil {
		errorMap := map[string]string{"error": db.Error.Error(), "code": "404"}
		errorJson, _ := json.Marshal(errorMap)
		w.Header().Set("Content-Type", "pkglication/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(errorJson)
		return
	}
	res, _ := json.Marshal(deletedPost)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
