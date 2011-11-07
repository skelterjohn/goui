package server

import (
	"http"
	"log"
)

func init() {
	log.Printf("goui listening on localhost:6061")
	go http.ListenAndServe(":6061", nil)
}
