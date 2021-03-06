package data

import (
	"encoding/json"
	"fmt"
	"io"
	"time"
)

// Product : defines the structure for an API Product
type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	SKU         string  `json:"sku"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

// Products : slice of (reference of) Product
type Products []*Product

// ToJSON : convert Products into JSON format
func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p) // encode myself (Products)
}

// FromJSON :
func (p *Product) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(p) // decode myself
}

// GetProducts : Get all products
func GetProducts() Products {
	return productList
}

// AddProduct : and new product
func AddProduct(p *Product) {
	p.ID = getNextID()
	productList = append(productList, p)
}

func getNextID() int {
	lp := productList[len(productList)-1]
	return lp.ID + 1
}

// UpdateProduct : update product based on id
func UpdateProduct(id int, newData *Product) error {
	_, pos, err := findProduct(id)

	if err != nil {
		return err
	}

	newData.ID = id
	productList[pos] = newData

	return nil
}

// ErrProductNotFound :
var ErrProductNotFound = fmt.Errorf("Product Not found")

func findProduct(id int) (*Product, int, error) {
	for position, p := range productList {
		if p.ID == id {
			return p, position, nil
		}
	}

	return nil, -1, ErrProductNotFound
}

var productList = []*Product{
	&Product{
		ID:          1,
		Name:        "Latte",
		Description: "Frothy milky coffee",
		Price:       2.45,
		SKU:         "abc123",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          2,
		Name:        "Espresso",
		Description: "Short and strong coffee without milk",
		Price:       1.99,
		SKU:         "def456",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}
