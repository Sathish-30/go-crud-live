package main

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	BookId     int    `json:"bookId"`
	BookName   string `json:"bookName"`
	AuthorName string `json:"authorName"`
	gorm.Model
}

func connect() {
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
}
