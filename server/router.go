package server

import (
	"net/http"

	"github.com/Tkdefender88/ButteAir/logger"

	"github.com/gorilla/mux"
)

func redirectHTTPS(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "https://"+r.Host+r.URL.String(),
		http.StatusMovedPermanently)
}

//NewRouters creates two mux routers from all the routes in the routes var above.
//One serves on http and the other https
func NewRouters() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	//file server for static assets
	fs := http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets")))
	router.PathPrefix("/assets/").Handler(logger.Logger(fs))

	//the routes that define our api
	router.Handle("/", http.HandlerFunc(Index)).Methods("GET")
	router.Handle("/airqual", http.HandlerFunc(AirQuality)).Methods("GET")
	router.Handle(
		"/data",
		isAuthorized(http.HandlerFunc(UpdateData)),
	).Methods("GET")

	return router
}
