package types

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type Post struct {
	ID      int    `json:"id" db:"id"`
	Title   string `json:"title" db:"title"`
	Content string `json:"content" db:"content"`
	Excerpt string `json:"excerpt" db:"excerpt"`
	Date    string `json:"date" db:"date"`
}

func NewLogger(servicename string) *log.Logger {
	timestamp := time.Now().Format("2006-01-02_15-04")
	file, err := os.OpenFile("bin/log/"+timestamp+"--"+servicename+".log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Dosya açılamadı: %v", err)
	}
	logger := log.New(file, servicename, log.Ldate|log.Ltime|log.Lshortfile)
	return logger
}

func GetEnv(key string) string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return os.Getenv(key)
}
