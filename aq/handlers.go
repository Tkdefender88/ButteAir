package aq

import (
	"net/http"

	"github.com/Tkdefender88/ButteAir/config"
)

type quality struct {
	Place    string
	Temp     int
	Humidity int
	Pm1      int
	Pm2      int
	Pm3      int
}

//Index serves index page
func Index(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(405), 405)
		return
	}

	qual := quality{
		"here",
		0,
		40,
		10000,
		200,
		300,
	}

	quals := []quality{
		qual,
		quality{
			"airport",
			20,
			40,
			110,
			210,
			310,
		},
		quality{
			"hospital",
			30,
			40,
			120,
			220,
			320,
		},
		quality{
			"Whale-Mart",
			30,
			40,
			120,
			220,
			320,
		},
		quality{
			"Sooubway",
			30,
			40,
			120,
			220,
			320,
		},
		quality{
			"Church",
			30,
			40,
			120,
			220,
			320,
		},
	}

	config.TPL.ExecuteTemplate(w, "airqual.html", quals)
}
