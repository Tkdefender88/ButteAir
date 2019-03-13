package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Tkdefender88/ButteAir/logger"
	"github.com/Tkdefender88/ButteAir/server"

	"github.com/rs/cors"
)

var (
	port = ":9000"
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
	fmt.Printf("Listening on port %s\n", port)
	fmt.Printf("Go to http://localhost%s to view\n", port)
	log.Fatal(http.ListenAndServe(port, c.Handler(loggedRouter)))
}
