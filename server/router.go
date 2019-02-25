package server

import (
	"log"
	"net/http"

	"github.com/Tkdefender88/ButteAir/aq"

	"github.com/Tkdefender88/ButteAir/logger"
	"github.com/gorilla/mux"
)

//Route describes a route for the page
type Route struct {
	Name       string
	Method     string
	Pattern    string
	HandleFunc http.HandlerFunc
}

//Routes is a collection of type Route
type Routes []Route

var fileServer http.HandlerFunc

func init() {
	//assert the file server is a handlerFunc
	fs, ok := http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))).(http.HandlerFunc)
	if !ok {
		log.Fatalln("Could not assert file server type")
	}

	fileServer = fs
}

var routes = Routes{
	Route{
		"index",
		http.MethodGet,
		"/",
		aq.Index,
	},
	Route{
		"airquality",
		http.MethodGet,
		"/airqual",
		aq.AirQuality,
	},
}

//NewRouter creates a mux router from all the routes in the routes var above.
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandleFunc
		handler = logger.Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
	return router
}
