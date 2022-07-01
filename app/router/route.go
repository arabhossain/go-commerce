package router

import (
	"github.com/gorilla/mux"
	"go-commerce/app/controllers/auth"
	"go-commerce/app/middlewares"
)

func SetupRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.PathPrefix("/v1")

	secure := router.PathPrefix("/secure").Subrouter()
	secure.Use(middlewares.JwtVerify)
	secure.HandleFunc("/register", auth.Login).Methods("GET")

	router.HandleFunc("/login", auth.Login).Methods("GET")
	//router.HandleFunc("/api/book/show/{id}", showBook).Methods("GET")

	return router
}
