package main

import (
	"log"
	"net/http"

	"github.com/Tkdefender88/ButteAir/aq"
)

var (
	config = oauth2.Config{
		ClientID:     "222222",
		ClientSecret: "222222",
		Scope:        []string{"all"},
		RedirectURL:  "http://localhost:9000/oauth2",
		Endpoint: oauth2.Endpoint{
			AuthURL:  "http://localhost:9000/authorize",
			TokenURL: "http://localhost:9000/token",
		},
	}
)

func main() {
	fs := http.FileServer(http.Dir("./assets"))
	http.HandleFunc("/", index)
	http.HandleFunc("/airqual", aq.Index)
	http.Handle("/assets/", http.StripPrefix("/assets", fs))

	log.Println("Listening ...")
	log.Fatal(http.ListenAndServe(":9000", nil))
}

func index(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/airqual", http.StatusSeeOther)
}
