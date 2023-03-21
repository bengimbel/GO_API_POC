package main

import (
	"log"
	"net/http"

	"github.com/bengimbel/go-bookstore/pkg/middlewares"
	"github.com/bengimbel/go-bookstore/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	userRoutes := r.PathPrefix("/user").Subrouter()
	bookRoutes := r.PathPrefix("/books").Subrouter()
	routes.RegisterUserRoutes(userRoutes)
	routes.RegisterBookStoreRoutes(bookRoutes)
	bookRoutes.Use(middlewares.AuthMiddleware)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
