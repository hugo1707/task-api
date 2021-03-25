package handler

import "net/http"

// Home handles all requests to /home
type Home struct {
}

// NewHome returns a new home handler ready to use
func NewHome() Home {
	return Home{}
}

func (h Home) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("home"))
}
