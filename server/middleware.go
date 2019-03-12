package server

import (
	"fmt"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
)

var (
	mySigningKey = []byte("captainjacksparrowsayshi")
)

func isAuthorized(endpoint http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Header["Token"] != nil {

			token, err := jwt.Parse(r.Header["Token"][0],
				func(token *jwt.Token) (interface{}, error) {
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf("There was an error")
					}
					return mySigningKey, nil
				},
			)
			if err != nil {
				fmt.Fprintf(w, err.Error())
			}

			if token.Valid {
				endpoint.ServeHTTP(w, r)
			}
		} else {
			fmt.Fprintf(w, "Not Authorized")
		}
	})
}
