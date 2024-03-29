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
	for _, book := range books {
		log.Printf("BookId %v BookName %v AuthorName %v", book.BookId, book.BookName, book.AuthorName)
	}
	w.WriteHeader(http.StatusAccepted)
	jsonData, err := json.Marshal(&response{Message: "Retrieved data"})
	if err != nil {
		log.Fatal("Error in parsing data")
	}
	w.Write(jsonData)
}

func addBookHandler(w http.ResponseWriter, r *http.Request) {
	var book Book
	json.NewDecoder(r.Body).Decode(&book)
	if result := db.Create(&book); result.Error != nil {
		log.Fatal("Failed to insert data to DB")
	}
	w.WriteHeader(http.StatusAccepted)
	jsonData, err := json.Marshal(&response{Message: "Added user"})
	if err != nil {
		log.Fatal("Error in parsing data")
	}
	w.Write(jsonData)
}
