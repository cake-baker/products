package main

import (
	"cakebakery.io/products/handlers"
	"flag"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	flag.BoolVar(&https, "https", false, "HTTPS server")
	flag.StringVar(&cert, "cert", "", "Cert file for HTTPS server")
	flag.StringVar(&key, "key", "", "Key file for HTTPS server")
	flag.Parse()

	router := mux.NewRouter()
	handlers.NewProductHandler(router).Init()

	if https {
		if err := http.ListenAndServeTLS(":8443", cert, key, router); err != nil {
			panic(err)
		}
	} else {
		if err := http.ListenAndServe(":8080", router); err != nil {
			panic(err)
		}
	}
}

var https bool
var cert string
var key string
