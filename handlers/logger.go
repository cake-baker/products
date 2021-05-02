package handlers

import (
	lg "log"
	"net/http"
)

func log(r *http.Request) {
	lg.Printf("%s %s", r.Method, r.RequestURI)
}
