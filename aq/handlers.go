package aq

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
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

// Index redirects to the airqual page
func Index(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/airqual", http.StatusSeeOther)
}

// AirQuality serves airquality page
func AirQuality(w http.ResponseWriter, r *http.Request) {
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

type airinfo struct {
	DeviceID string         `json:"deviceID"`
	Temp     string         `json:"temp"`
	Humidity string         `json:"humidity"`
	PM10     string         `json:"pm1"`
	PM25     string         `json:"pm25"`
	PM100    string         `json:"pm100"`
	Location location       `json:"location"`
	Time     collectionTime `json:"now"`
}

type collectionTime struct {
	Time string `json:"time"`
	Date string `json:"date"`
}

type location struct {
	Lat  float64 `json:"lat"`
	Long float64 `json:"long"`
}

// UpdateData is where POST requests are recieved from the dragino radio with
// new data
func UpdateData(w http.ResponseWriter, r *http.Request) {
	var info airinfo

	err := json.NewDecoder(r.Body).Decode(&info)
	if err != nil {
		if err == io.EOF {
			log.Println("Empty request body")
			w.WriteHeader(http.StatusBadRequest)
		} else {
			log.Println("Error: ", err.Error())
			w.WriteHeader(422) //unprocessable entity
			if err := json.NewEncoder(w).Encode(err); err != nil {
				log.Println("Failed encoding error into response")
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			return
		}
	}

	fmt.Fprintf(w, info.Humidity)
}
