package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := http.NewServeMux()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hellooooooooOOOOOOOOOOOOOOOOOOOOOOOOOOOOoooo")
	})

	log.Println("server started")
	http.ListenAndServe(":8080", r)
	chiR := chi.NewRouter
	_ = chiR
}
