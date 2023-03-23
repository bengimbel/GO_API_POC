package routes

// absolute paths
import (
	"github.com/bengimbel/go-bookstore/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterPostRoutes = func(router *mux.Router) {
	router.HandleFunc("/", controllers.GetPosts).Methods("GET")
	router.HandleFunc("/", controllers.CreatePost).Methods("POST")
	router.HandleFunc("/{postId}", controllers.GetPostById).Methods("GET")
	router.HandleFunc("/{postId}", controllers.UpdatePost).Methods("PUT")
	router.HandleFunc("/{postId}", controllers.UpdatePost).Methods("PATCH")
	router.HandleFunc("/{postId}", controllers.DeletePost).Methods("DELETE")
	router.HandleFunc("/{postId}/comment", controllers.CreateComment).Methods("POST")
	router.HandleFunc("/{postId}/like", controllers.CreateLike).Methods("POST")
}
