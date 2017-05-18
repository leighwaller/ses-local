package main

import (
	"net/http"
	"log"
	"os"
	"ses-local/route"
)

func main() {
	port := serverPort()
	router := route.CreateRouter()
	log.Printf("Started SES service on port: %s", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

func serverPort() string {
	port := os.Getenv("ses.server.port")
	if port == "" {
		return "8080"
	}
	return port
}

