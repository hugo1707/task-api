package handler

import (
	"log"
	"net/http"

	"github.com/hugo1707/task-api/data"
)

// Product handles all requests to /products
type Product struct {
	l *log.Logger
}

// NewProduct returns a new product handler ready to use
func NewProduct(l *log.Logger) Product {
	return Product{l}
}

func (p Product) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		p.get(rw, r)
	} else {
		rw.Header().Set("reason", "try again")
		http.Error(rw, "method not supported", http.StatusMethodNotAllowed)
		return 
	}
}

func (p Product) get(rw http.ResponseWriter, r *http.Request) {
	products := data.ReadAll()
	err := products.WriteJSON(rw)
	if err != nil {
		http.Error(rw, "internal server error", http.StatusInternalServerError)
		return 
	}
} 
