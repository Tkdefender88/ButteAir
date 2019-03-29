package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Tkdefender88/ButteAir/logger"

	"github.com/gorilla/mux"
)

func redirectHTTPS(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "https://"+r.Host+r.URL.String(),
		http.StatusMovedPermanently)
}

// NewHTTPRouter creates a router for http that redirects back to https
func NewHTTPRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	router.Path("/").Handler(http.HandlerFunc(redirectHTTPS))

	return router
}

// MakeServer given a mux router will make a server using the router as
// the handler
func MakeServer(mux http.Handler) *http.Server {
	//set timeouts so malicious or slow clients don't hog resources
	return &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      mux,
	}
}

// RedirectServer makes a server that redirects to https
func RedirectServer() *http.Server {
	handleRedirect := func(w http.ResponseWriter, r *http.Request) {
		newURI := fmt.Sprintf("https://%s%s", r.Host, r.URL)
		http.Redirect(w, r, newURI, http.StatusMovedPermanently)
	}

	router := mux.NewRouter().StrictSlash(true)
	router.PathPrefix("/").HandlerFunc(handleRedirect)

	return MakeServer(router)
}

// NewRouter creates a mux router to serve https.
func NewRouter() *mux.Router {
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
