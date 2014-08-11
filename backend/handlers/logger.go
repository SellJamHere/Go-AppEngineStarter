package handlers

import (
	"appengine"
	"net/http"
	"time"

	"github.com/codegangsta/martini"
)

const (
	logFormat = `[api] rem_ip="%s"        rcvd_cookies="%s"        url="%s %s"        status="%d"        referer="%s"        user_agent="%s"        content_length="%d"        content_type="%s"        form="%s"        response_time="%s"`
)

type Time interface {
	GetStartTime() time.Time
}

type timer struct {
	startTime time.Time
}

func (t *timer) GetStartTime() time.Time {
	return t.startTime
}

func Logger() martini.Handler {
	return func(w http.ResponseWriter, r *http.Request, c martini.Context) {
		start := time.Now()

		c.MapTo(&timer{start}, (*Time)(nil))

		rw := w.(martini.ResponseWriter)
		c.Next()

		gaeContext := appengine.NewContext(r)
		gaeContext.Infof(logFormat+"\n", r.Header.Get("X-Forwarded-For"), r.Cookies(), r.Method, r.URL, rw.Status(), r.Referer(), r.UserAgent(), r.ContentLength, r.Header.Get("Content-Type"), r.Form, time.Now().Sub(start))
	}
	return func(res http.ResponseWriter, c martini.Context) {

	}
}
