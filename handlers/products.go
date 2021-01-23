package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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

// GetProducts : fetch all products
func (p *Product) GetProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET Products")

	// fetch the products from datastore
	lp := data.GetProducts()

	// serialize the list to JSON
	err := lp.ToJSON(rw) // d, err := json.Marshal(lp)

	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
		return
	}
}

// AddProduct : add a new product
func (p *Product) AddProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST Products")

	prod := &data.Product{}

	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
	}

	p.l.Printf("Prod : %#v\n", prod)
	data.AddProduct(prod)
}

// UpdateProducts : update product
func (p *Product) UpdateProducts(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		http.Error(rw, "id should be integer", http.StatusBadRequest)
	}

	p.l.Println("Handle PUT Products", id)

	prod := &data.Product{}
	p.l.Println("prod (before) : ", prod)

	err = prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
		return
	}

	p.l.Println("prod (after) : ", prod)

	err = data.UpdateProduct(id, prod)

	if err == data.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
}
