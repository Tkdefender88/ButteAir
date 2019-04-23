package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/user"
	"path/filepath"

	"github.com/Tkdefender88/ButteAir/logger"
	"github.com/Tkdefender88/ButteAir/server"

	"github.com/rs/cors"
)

const (
	domain = "justinbak.com"
)

func main() {

	//create router
	router := server.NewRouter()
	loggedRouter := logger.Logger(router)

	var httpsSrv *http.Server
	//var m *autocert.Manager

	//set up CORS middleware
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "PUT", "POST", "OPTIONS"},
	})
	/*
		//set up cert manager
		hostPolicy := func(ctx context.Context, host string) error {
			allowedHost := domain
			if host == allowedHost {
				return nil
			}
			return fmt.Errorf("acme/autocert: only %s host is allowed",
				allowedHost)
		}

		m = &autocert.Manager{
			Prompt:     autocert.AcceptTOS,
			HostPolicy: hostPolicy,
		}

		dir, err := cacheDir()
		if err == nil {
			m.Cache = autocert.DirCache(dir)
		}
	*/
	//create the server with the routers
	httpsSrv = server.MakeServer(c.Handler(loggedRouter))
	httpsSrv.Addr = ":8080"
	//httpsSrv.TLSConfig = &tls.Config{GetCertificate: m.GetCertificate}

	fmt.Printf("Starting http/https server on %s\n", httpsSrv.Addr)

	//serve https
	log.Fatal(httpsSrv.ListenAndServe())
}

// cacheDir creates a consistent cache directory for the tls certificates
func cacheDir() (string, error) {
	var dir string
	if u, _ := user.Current(); u != nil {
		dir = filepath.Join(os.TempDir(), "cert-cache-autocert-"+u.Username)
		if err := os.MkdirAll(dir, 0700); err != nil {
			return "", err
		}
	}
	return dir, nil
}
