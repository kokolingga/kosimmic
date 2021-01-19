package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		d, err := ioutil.ReadAll(r.Body) // read the body

		if err != nil {
			log.Println("Error reading body", err)
			http.Error(rw, "Unable to read request body", http.StatusBadRequest)
			return
		}

		log.Printf("Hello %s\n", d)
		fmt.Fprintf(rw, "Hello %s\n", d) // write the response
	})

	http.HandleFunc("/goodbye", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("Goodbye!")
		fmt.Fprintln(rw, "Goodbye!")
	})

	log.Println("Starting Server")
	err := http.ListenAndServe(":9090", nil)
	log.Fatal(err)
}
