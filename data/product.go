package data

import (
	"encoding/json"
	"io"
)

// Product defines the representation
type Product struct {
	Name string `json:"name"`
	Price float64 `json:"price"`
	Category string `json:"category"`
}

// Products is a list of products
type Products []Product

// WriteJSON writes the JSON representation of the list into `w`
func (p Products) WriteJSON(w io.Writer) error {
	return json.NewEncoder(w).Encode(p)
}

// ReadAll returns all products
func ReadAll() Products {
	return products
}

var products = Products{
	{
		Name: "lapis",
		Price: 1.99,
		Category: "papelaria",
	},
	{
		Name: "Cavelete Piso",
		Price: 56,
		Category: "utilidades",
	},
}