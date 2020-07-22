package main

import (
	"log"
	"net/http"
	"os"

	"github.com/ArchieSpinos/eCommerce/controllers"
	"github.com/gorilla/handlers"
)

func main() {
	// mongo.InitMongo()

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := controllers.NewRouter()

	allowdOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET"})

	log.Fatal(http.ListenAndServe(":"+port, handlers.CORS(allowdOrigins, allowedMethods)(router)))
}
