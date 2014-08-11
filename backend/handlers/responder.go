package handlers

import (
	"encoding/json"
	"net/http"

	"backend/errors"

	"github.com/codegangsta/martini"
)

const (
	ContentType      = "Content-Type"
	ContentLength    = "Content-Length"
	ContentLocation  = "Location"
	ContentJSON      = "application/json; charset=UTF-8"
	ContentTextPlain = "text/plain; charset=UTF-8"
)

type Respond interface {
	Valid(status int, v interface{})
	ValidRedirect(status int, url string, req *http.Request)
	Error(error)
}

func Responder() martini.Handler {
	return func(res http.ResponseWriter, c martini.Context) {
		c.MapTo(&responder{res}, (*Respond)(nil))
	}
}

type responder struct {
	http.ResponseWriter
}

func (r *responder) Valid(status int, v interface{}) {
	if v == nil {
		r.WriteHeader(status)
		return
	}

	result, err := json.Marshal(v)
	if err != nil {
		http.Error(r, err.Error(), 500)
		return
	}

	// json rendered fine, write out the result
	r.Header().Set(ContentType, ContentJSON)
	r.WriteHeader(status)
	r.Write(result)
}

func (r *responder) ValidRedirect(status int, url string, req *http.Request) {
	http.Redirect(r, req, url, status)
}

func (r *responder) Error(err error) {
	serverErr, ok := err.(*errors.ServerError)
	if !ok {
		serverErr = errors.New(err, "", 500)
	}

	result, jsonErr := json.MarshalIndent(serverErr, "  ", "")
	if jsonErr != nil {
		http.Error(r, jsonErr.Error(), 500)
		return
	}

	// json rendered fine, write out the result
	r.Header().Set(ContentType, ContentJSON)
	r.WriteHeader(serverErr.StatusCode)
	r.Write(result)
}
