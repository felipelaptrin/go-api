package routes

import (
	"github.com/felipelaptrin/go-api/pkg/controllers"
	"github.com/gorilla/mux"
)

func RegisterUserRoutes(router *mux.Router) {
	router.HandleFunc("/user", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/user", controllers.GetAllUsers).Methods("GET")
	router.HandleFunc("/user/{id}", controllers.GetUserById).Methods("GET")
	router.HandleFunc("/user/{id}", controllers.UpdateUserById).Methods("PUT")
	router.HandleFunc("/user/{id}", controllers.DeleteUserById).Methods("DELETE")
}
