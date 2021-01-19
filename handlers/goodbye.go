package handlers

import (
	"log"
	"net/http"
)

// Goodbye : Handler (for Goodbye) that satisfies http Handler interfaces
type Goodbye struct {
	l *log.Logger
}

// NewGoodbye : idiomatic principle of creating go code
func NewGoodbye(l *log.Logger) *Goodbye {
	return &Goodbye{l}
}

func (g *Goodbye) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	g.l.Println("Goodbye!")
	rw.Write([]byte("Goodbye!\n"))
}
