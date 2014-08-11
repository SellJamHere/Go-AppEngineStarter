package controllers

import (
	"net/http"

	"backend/handlers"
)

func init() {
	m.Get("/sample", getSample)
}

func getSample(r handlers.Respond, req *http.Request) {
	r.Valid(200, "sample")
}
