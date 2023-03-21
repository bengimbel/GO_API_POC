package routes

// absolute paths
import (
	"github.com/bengimbel/go-bookstore/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/", controllers.GetBook).Methods("GET")
	router.HandleFunc("/{bookId}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/{bookId}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/{bookId}", controllers.UpdateBook).Methods("PATCH")
	router.HandleFunc("/{bookId}", controllers.DeleteBook).Methods("DELETE")
}
