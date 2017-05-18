package handler

import (
	"log"
	"net/http"
	"time"
)

func RequestLoggingAdapter(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer timer(time.Now(), r, name)
		inner.ServeHTTP(w, r)
	})
}

func timer(start time.Time, r *http.Request, name string) {
	elapsed := time.Since(start)
	log.Printf("%s  %s  %s  %s", r.Method, r.RequestURI, name, elapsed)
}
