package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Failed to load env")
	}
	connect()
}

func main() {
	const PORT string = ":8080"
	server := http.NewServeMux()
	server.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]string{
			"msg": "In Home route",
		})
	})
	server.HandleFunc("GET /getBooks", getAllBooksHandler)
	server.HandleFunc("GET /getBook/{id}", getBookHandler)
	server.HandleFunc("POST /addBook", addBookHandler)
	http.ListenAndServe(PORT, respJsonMiddleware(server))
}

func respJsonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("The endpoint of the request is %v and the http request method is %v", r.URL, r.Method)
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
