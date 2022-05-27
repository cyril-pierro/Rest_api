package router

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func StartRoutes() {
	r := mux.NewRouter()
	r.HandleFunc("/api/books", GetBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", GetBook).Methods("GET")
	r.HandleFunc("/api/books", AddBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", UpdateBooks).Methods("PUT")
	r.HandleFunc("/api/books/{id}", DeleteBooks).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":9000", r))
}
