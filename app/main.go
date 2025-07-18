package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	r := http.NewServeMux()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "helloooooooooooo")
	})

	log.Println("server started")
	http.ListenAndServe(":8080", r)
}
