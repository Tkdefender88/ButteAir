package aq

import (
	"net/http"

	"github.com/Tkdefender88/ButteAir/config"
)

type quality struct {
	Temp     int
	Humidity int
	Pm1      int
	Pm2      int
	Pm3      int
}

//Index redirects to the airqual page
func Index(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/airqual", http.StatusSeeOther)
}

//AirQuality serves airquality page
func AirQuality(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(405), 405)
		return
	}

	qual := quality{
		0,
		40,
		100,
		200,
		300,
	}

	config.TPL.ExecuteTemplate(w, "airqual.gohtml", qual)
}
