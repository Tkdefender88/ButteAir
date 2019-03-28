package main

import (
	"context"
	"crypto/tls"
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"golang.org/x/crypto/acme/autocert"

	"github.com/Tkdefender88/ButteAir/logger"
	"github.com/Tkdefender88/ButteAir/server"

	"github.com/rs/cors"
)

var (
	port       = ":9000"
	httpsport  = ":443"
	production *bool
)

const (
	certsDir = "/app/certs/"
)

func makeHTTPServer() *http.Server {
	mux := &http.ServeMux{}
	mux.HandleFunc("/", server.Index)

	//set timeouts so that slow or malicious clients don't hog resources
	return &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		IdleTimeout:  120 * time.Second,
	}
}

func main() {
	production = flag.Bool("prod", false, "Run the server in production?")
	flag.Parse()

	var httpsSrv *http.Server
	var m *autocert.Handler

	if *production {

		hostPolicy := func(ctx context.Context, host string) error {
			allowedHost := "justinbak.com"
			if host == allowedHost {
				return nil
			}
			return fmt.Errorf("acme/autocert: only %s host is allowed",
				allowedHost)
		}

		httpsSrv = makeHTTPServer()
		m := &autocert.Manager{
			Prompt:     autocert.AcceptTOS,
			HostPolicy: hostPolicy,
			Cache:      autocert.DirCache(certsDir),
		}
		httpsSrv.Addr = ":443"
		httpsSrv.TLSConfig = &tls.Config{GetCertificate: m.GetCertificate}

		go func() {
			err := httpsSrv.ListenAndServeTLS("", "")
			if err != nil {
				log.Fatalf("httpsSrv.ListenAndServeTLS() failed with: %s", err)
			}
		}()
	}

	httpSrv := makeHTTPServer()
	if m != nil {
		httpSrv.Handler = m.HTTPHandler(httpSrv.Handler)
	}

	router, httpsrouter := server.NewRouters()
	loggedRouter := logger.Logger(router)
	loggedHTTPSRouter := logger.Logger(httpsrouter)

	//set up CORS middleware
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "PUT", "OPTIONS"},
	})

	//Start the sever
	fmt.Printf("Listening on port %s\n", port)
	fmt.Printf("Go to http://localhost%s to view\n", port)

	go log.Fatal(
		http.ListenAndServeTLS(
			httpsport,
			"cert.pem",
			"key.pem",
			c.Handler(loggedHTTPSRouter),
		),
	)
	log.Fatal(http.ListenAndServe(port, c.Handler(loggedRouter)))
}
