package routes

// absolute paths
import (
	"github.com/bengimbel/go-bookstore/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterUserRoutes = func(router *mux.Router) {
	router.HandleFunc("/register", controllers.RegisterUser).Methods("POST")
	router.HandleFunc("/login", controllers.GenerateToken).Methods("POST")
}
