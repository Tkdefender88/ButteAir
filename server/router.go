package server

import (
	"net/http"

	"github.com/Tkdefender88/ButteAir/logger"

	"github.com/gorilla/mux"
)

func redirectHTTPS(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "https://127.0.0.1:443"+r.RequestURI,
		http.StatusMovedPermanently)
}

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
		http.HandlerFunc(redirectHTTPS),
	)

	return router
}
