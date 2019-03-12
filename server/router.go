package server

import (
	"net/http"

	"github.com/Tkdefender88/ButteAir/logger"

	"github.com/gorilla/mux"
)

//NewRouter creates a mux router from all the routes in the routes var above.
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	//file server for static assets
	fs := http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets")))
	router.PathPrefix("/assets/").Handler(logger.Logger(fs))

	//the routes that define our api
	router.Handle(
		"/",
		http.HandlerFunc(Index),
	)
	router.Handle(
		"/airqual",
		http.HandlerFunc(AirQuality),
	)
	router.Handle(
		"/data",
		isAuthorized(
			http.HandlerFunc(UpdateData),
		),
	)

	return router
}
