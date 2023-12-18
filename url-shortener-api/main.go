package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"url-shortener-api/db"
	"url-shortener-api/http"
)

func init() {
	log.Println("loading .env file ...")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
	log.Println(".env file loaded successfully")

	log.Println("establishing connection to database ...")
	dbConfig := db.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Pass:     os.Getenv("DB_PASS"),
		Database: os.Getenv("DB_DB"),
	}
	fmt.Println(dbConfig)
	db.Init(dbConfig)
	log.Println("successfully established connection to database")
}

func main() {
	serverConfig := http.Config{Host: os.Getenv("HOST"), Port: os.Getenv("PORT")}
	server := http.NewServer(serverConfig)
	server.Run()
}
