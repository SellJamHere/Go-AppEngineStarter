package controllers

import (
	"fmt"
	"net/http"
	"time"

	"backend/handlers"

	"github.com/codegangsta/martini"
	"github.com/dustin/go-humanize"
)

var startTime time.Time
var m = newMainHandler()

func init() {
	// Inject into parent handler.
	http.Handle("/", m)

	m.Get("/", index)
}

func newMainHandler() *martini.ClassicMartini {
	startTime = time.Now()

	r := martini.NewRouter()
	m := martini.New()
	m.Action(r.Handle)

	// Middleware
	m.Use(handlers.Logger())
	m.Use(martini.Recovery())
	m.Use(handlers.Responder())
	m.Use(handlers.Origin(nil))

	return &martini.ClassicMartini{m, r}
}

func index(r handlers.Respond, t handlers.Time) {
	ok := map[string]string{
		"status":       "Up",
		"responseTime": fmt.Sprintf("%s", time.Now().Sub(t.GetStartTime())),
		"downTime":     fmt.Sprintf("%s", humanize.Time(startTime)),
	}

	r.Valid(200, ok)
}
