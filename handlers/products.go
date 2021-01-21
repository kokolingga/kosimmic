package handlers

import (
	"log"
	"net/http"
	"regexp"
	"strconv"

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
	// handle the request for a list of products
	if r.Method == http.MethodGet {
		p.getProducts(rw, r)
		return
	}

	if r.Method == http.MethodPost {
		p.addProduct(rw, r)
		return
	}

	if r.Method == http.MethodPut {
		p.l.Println("PUT", r.URL.Path)

		// expect the id in the URI
		reg := regexp.MustCompile(`/([0-9]+)`)
		g := reg.FindAllStringSubmatch(r.URL.Path, -1)

		p.l.Println("g", g)

		if len(g) != 1 {
			p.l.Printf("Invalid URI : more than one id (len : %v)", len(g))
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}

		if len(g[0]) != 2 {
			p.l.Println("Invalid URI : more than capture group", len(g[0]))
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}

		idString := g[0][1]
		id, err := strconv.Atoi(idString)
		if err != nil {
			p.l.Println("Invalid URI : unable to convert to number", idString)
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
		}

		// p.l.Println("Got id", id)

		p.updateProducts(id, rw, r)
	}

	// catch all
	// if no method is satisfied return an error
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Product) getProducts(rw http.ResponseWriter, r *http.Request) {
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

func (p *Product) addProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST Products")

	prod := &data.Product{}

	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
	}

	p.l.Printf("Prod : %#v\n", prod)
	data.AddProduct(prod)
}

func (p *Product) updateProducts(id int, rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle PUT Products")

	prod := &data.Product{}
	p.l.Println("prod (before) : ", prod)

	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
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
