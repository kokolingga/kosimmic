package main

import (
	"log"
	"net/http"
	"os"

	"github.com/kokolingga/kosimmic/handlers"
)

func main() {
	l := log.New(os.Stdout, "[product-api] ", log.LstdFlags)

	hh := handlers.NewHello(l)   // hello handler
	gh := handlers.NewGoodbye(l) // goodbye handler

	sm := http.NewServeMux()
	sm.Handle("/", hh)
	sm.Handle("/goodbye", gh)

	http.ListenAndServe(":9090", sm)
}
