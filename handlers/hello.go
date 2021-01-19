package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Hello : Handler (for Hello) that satisfies http Handler interfaces
type Hello struct {
	l *log.Logger
}

// NewHello : idiomatic principle of creating go code
func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.l.Println("Hello World")
	d, err := ioutil.ReadAll(r.Body) // read the body

	if err != nil {
		http.Error(rw, "Ooppss", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(rw, "Hello %s\n", d) // write the response
}
