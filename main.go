package main

import (
	"log"
	"net/http"

	"github.com/Tkdefender88/joelSite/aq"
)

func main() {
	fs := http.FileServer(http.Dir("./assets"))
	http.HandleFunc("/", index)
	http.HandleFunc("/airqual", aq.Index)
	http.Handle("/resources/", http.StripPrefix("/resources", fs))

	log.Println("Listening ...")
	log.Fatal(http.ListenAndServe(":9000", nil))
}

func index(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/airqual", http.StatusSeeOther)
}
