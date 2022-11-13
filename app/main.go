package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", hello).Methods(http.MethodGet)

	http.ListenAndServe(":8080", r)
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello")
}
