package router

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/cyril_pierro/Rest_api_2/helpers"
	"github.com/cyril_pierro/Rest_api_2/model"
	"github.com/gorilla/mux"
)

func GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var allBooks []model.Book
	allBooks = helpers.GetAllBooks(allBooks)
	json.NewEncoder(w).Encode(allBooks)
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	param := mux.Vars(r)
	id := helpers.ParseInt(param["id"])
	book := helpers.GetBookWithId(id)
	json.NewEncoder(w).Encode(&book)
}

func AddBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book model.Book
	err := json.NewDecoder(r.Body).Decode(&book)
	helpers.CheckError(err)
	fmt.Println(book)
	helpers.AddBook(&book)
	json.NewEncoder(w).Encode(&book)
	defer r.Body.Close()
}

func UpdateBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	param := mux.Vars(r)
	id := helpers.ParseInt(param["id"])
	book := helpers.GetBookWithId(id)
	var new_book model.Book
	json.NewDecoder(r.Body).Decode(&book)
	helpers.UpdateBook(&book, &new_book)
	json.NewEncoder(w).Encode(&new_book)
	defer r.Body.Close()
}

func DeleteBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	param := mux.Vars(r)
	id := helpers.ParseInt(param["id"])
	book := helpers.GetBookWithId(id)
	helpers.DeleteBook(&book, id)
	json.NewEncoder(w).Encode(&model.Message{Message: "Book has been deleted"})
}
