package routes

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	controller "test-prepare/controller/auth"
	"test-prepare/middleware"
)

func NewRouter() {
	r := mux.NewRouter()

	auth := r.PathPrefix("/auth").Subrouter()
	auth.HandleFunc("/register", controller.Register).Methods("POST")
	auth.HandleFunc("/login", controller.Login).Methods("POST")

	api := auth.PathPrefix("/api").Subrouter()
	api.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("/")
	}).Methods("GET")
	api.HandleFunc("/products", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("/products")
	}).Methods("GET")
	api.HandleFunc("/articles", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("/articles")
	}).Methods("GET")
	api.Use(middleware.JWTMiddleware)

	err := http.ListenAndServe(":7070", r)
	if err != nil {
		log.Fatal(err)
	}

}
