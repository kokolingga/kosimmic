package handlers

import (
	"log"
	"net/http"

	"github.com/kokolingga/kosimmic/data"
)

// Product : Handler (for Product) that satisfies http Handler interfaces
type Product struct {
	l *log.Logger
}

// NewProducts : idiomatic principle of creating go code
func NewProducts(l *log.Logger) *Product {
	return &Product{l}
}

func (p *Product) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts()
	err := lp.ToJSON(rw) // d, err := json.Marshal(lp)

	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
		return
	}
}
