package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type response struct {
	Message string `json:"message"`
}

func getAllBooksHandler(w http.ResponseWriter, r *http.Request) {
	var books []Book
	if result := db.Find(&books); result.Error != nil {
		log.Fatal("Error in retrieving data")
	}
	w.WriteHeader(http.StatusAccepted)
	err := json.NewEncoder(w).Encode(&books)
	if err != nil {
		log.Fatal("Error in parsing data")
	}
}

func getBookHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	var res Book
	if result := db.First(&res, "id = ?", id); result.Error != nil {
		w.WriteHeader(http.StatusNoContent)
		w.Write([]byte("Book not found"))
	}
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(&res); err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
	}
}

func addBookHandler(w http.ResponseWriter, r *http.Request) {
	var book Book
	json.NewDecoder(r.Body).Decode(&book)
	if result := db.Create(&book); result.Error != nil {
		log.Fatal("Failed to insert data to DB")
	}
	w.WriteHeader(http.StatusAccepted)
	err := json.NewEncoder(w).Encode(&response{Message: "Added user to database"})
	if err != nil {
		log.Fatal("Error in parsing data")
	}
}
