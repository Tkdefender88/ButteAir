package aq

import (
	"net/http"

	"github.com/Tkdefender88/joelSite/config"
)

//Index serves index page
func Index(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(405), 405)
		return
	}

	qual := "45%"

	config.TPL.ExecuteTemplate(w, "airqual.gohtml", qual)
}
