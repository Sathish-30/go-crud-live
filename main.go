package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Book struct {
	BookId     int    `json:"bookId"`
	BookName   string `json:"bookName"`
	AuthorName string `json:"authorName"`
	gorm.Model
}

var db *gorm.DB

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Failed to load env")
	}
}

func main() {
	const PORT string = ":8080"
	server := http.NewServeMux()
	dsn := os.Getenv("DATABASE_URL")
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err := db.AutoMigrate(&Book{}); err != nil {
		log.Fatal("Failed to migrate database")
	}
	if err != nil {
		log.Fatal("Failed to connect to database")
	}
	log.Println("Database connected")
	server.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]string{
			"msg": "In Home route",
		})
	})
	server.HandleFunc("GET /getBooks", getAllBooksHandler)
	http.ListenAndServe(PORT, server)
}

func getAllBooksHandler(w http.ResponseWriter, r *http.Request) {

}
