package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		log.Printf("method: %s", r.Method)
		log.Printf("path: %s", r.URL.Path)

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(rw, "someting went wrong", http.StatusBadRequest)
			return
		}
		log.Printf("body: %s", body)

		fmt.Fprintf(rw, "requested with\nmethod: %s\npath: %s\nbody: %s\n", r.Method, r.URL.Path, body)
	})

	http.HandleFunc("/admin", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("admin route /admin:", r)
	})

	http.ListenAndServe(":3000", nil)
}
