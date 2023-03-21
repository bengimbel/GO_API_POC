package routes

// absolute paths
import (
	"github.com/bengimbel/go-bookstore/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterUserRoutes = func(router *mux.Router) {
	// router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	// router.HandleFunc("/book/", controllers.GetBook).Methods("GET")
	// router.HandleFunc("/book/{bookId}", router.Use(middleware(controllers.GetBookById))).Methods("GET")
	// router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")
	// router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PATCH")
	// router.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")
	router.HandleFunc("/register", controllers.RegisterUser).Methods("POST")
	router.HandleFunc("/login", controllers.GenerateToken).Methods("POST")
}
