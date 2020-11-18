package main

import (
	"log"
	"net/http"

	"urlShortener/web"
)

func main() {
	mux := http.NewServeMux()
	web.RegisterHandlers(mux)

	if err := http.ListenAndServe(":6985", mux); err != nil {
		log.Fatalf("unable to start server: %v\n", err)
	}
}
