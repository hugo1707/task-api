package middleware

import (
	"net/http"
	"time"
)

// Slow the execution handler
func Slow(h http.Handler, duration time.Duration) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		time.Sleep(duration)
		h.ServeHTTP(rw, r)
	})
}
