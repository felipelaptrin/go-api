package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/felipelaptrin/go-api/pkg/config"
	"github.com/felipelaptrin/go-api/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("@@@@@@@@@@@@@@@@@@@@@@@@@@")

	// DB Setup
	config.Connect()
	config.Migrate()

	// Routes Setup
	router := mux.NewRouter()
	routes.RegisterUserRoutes(router)

	// Server Start
	log.Println("Starting server.")
	log.Fatal(http.ListenAndServe(":8080", router))
}
