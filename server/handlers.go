package server

import (
	"encoding/json"
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

//Airinfo represents a json object that would be recieved from the device about
//an update to the airquality at a location
type Airinfo struct {
	DeviceID string         `json:"deviceID"`
	Name     string         `json:"name"`
	Temp     string         `json:"temp"`
	Humidity string         `json:"humidity"`
	PM10     string         `json:"pm10"`
	PM25     string         `json:"pm25"`
	PM1      string         `json:"pm1"`
	Location Location       `json:"location"`
	Time     CollectionTime `json:"now"`
}

//CollectionTime is the time and date when a sample was collected
type CollectionTime struct {
	Time string `json:"time"`
	Date string `json:"date"`
}

//Location is a representation of where each sampling device is placed
//represented by latitude and longitude
type Location struct {
	Lat  float64 `json:"lat"`
	Long float64 `json:"long"`
}

var (
	devices = make(map[string]*Airinfo)
)

// Index redirects to the airqual page
func Index(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/airqual", http.StatusSeeOther)
}

//NotImplemented is a hanlder to use in place of a final handler while defining
//routes and the api of the application
func NotImplemented(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Not yet Implemented"))
}

// AirQuality serves airquality page
func AirQuality(w http.ResponseWriter, r *http.Request) {
	config.TPL.ExecuteTemplate(w, "airqual.html", devices)
}

// UpdateData is where POST requests are recieved from the dragino radio with
// new data
func UpdateData(w http.ResponseWriter, r *http.Request) {
	var info Airinfo

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

	devices[info.Name] = &info
}
