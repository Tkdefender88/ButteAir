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
	port      = ":9000"
	httpsport = ":443"
)

func main() {
	router, httpsrouter := server.NewRouters()
	loggedRouter := logger.Logger(router)
	loggedHTTPSRouter := logger.Logger(httpsrouter)

	//set up CORS middleware
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "PUT", "OPTIONS"},
	})

	//Start the sever
	fmt.Printf("Listening on port %s\n", port)
	fmt.Printf("Go to http://localhost%s to view\n", port)

	go log.Fatal(
		http.ListenAndServeTLS(
			httpsport,
			"cert.pem",
			"key.pem",
			c.Handler(loggedHTTPSRouter),
		),
	)
	log.Fatal(http.ListenAndServe(port, c.Handler(loggedRouter)))
}
