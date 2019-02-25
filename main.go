package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Tkdefender88/ButteAir/server"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/rs/cors"
)

var (
	/*	config = oauth2.Config{
			ClientID:     "222222",
			ClientSecret: "222222",
			Scopes:       []string{"all"},
			RedirectURL:  "http://localhost:9000/oauth2",
			Endpoint: oauth2.Endpoint{
				AuthURL:  "http://localhost:9000/authorize",
				TokenURL: "http://localhost:9000/token",
			},
		}
	*/
	mySigningKey = []byte("captainjacksparrowsayshi")
)

func main() {
	router := server.NewRouter()

	//Add the file server to the new router
	fs := http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets")))
	router.PathPrefix("/assets/").Handler(fs)

	//set up CORS
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "OPTIONS"},
	})

	//Start the sever
	log.Println("Listening ...")
	log.Fatal(http.ListenAndServe(":9000", c.Handler(router)))
}

func update(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(405), 405)
	}
}

func isAuthorized(endpoint func(http.ResponseWriter, *http.Request)) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Header["Token"] != nil {

			token, err := jwt.Parse(r.Header["Token"][0],
				func(token *jwt.Token) (interface{}, error) {
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf("There was an error")
					}
					return mySigningKey, nil
				})

			if err != nil {
				fmt.Fprintf(w, err.Error())
			}

			if token.Valid {
				endpoint(w, r)
			}
		} else {

			fmt.Fprintf(w, "Not Authorized")
		}
	})
}
