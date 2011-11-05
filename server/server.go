package server

import (
	"http"
)

func init() {
	go http.ListenAndServe(":6061", nil)
}
