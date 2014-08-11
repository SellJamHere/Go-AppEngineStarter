package main

import (
	"net/http"
)

//App Engine implements its own main function, so init is used instead
func init() {

	/*
	   Remove this if you don't build an API Server
	*/
	// No Favicon this is an API Server. But god forbid you actually use this in a browser.
	http.Handle("/favicon.ico", http.NotFoundHandler())
}
