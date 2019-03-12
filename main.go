package main

import (
	"log"
	"net/http"

	"github.com/Tkdefender88/ButteAir/logger"
	"github.com/Tkdefender88/ButteAir/server"

	"github.com/rs/cors"
)

func main() {
	router := server.NewRouter()
	loggedRouter := logger.Logger(router)

	//set up CORS middleware
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "PUT", "OPTIONS"},
	})

	//Start the sever
	log.Println("Listening ...")
	log.Fatal(http.ListenAndServe(":9000", c.Handler(loggedRouter)))
}
