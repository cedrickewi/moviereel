package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/cedrickewi/moviereel/handlers"
	"github.com/cedrickewi/moviereel/logger"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq" //import postges driver for database connection 
)

func initializeLogger() logger.Logger {
	logint, err := logger.NewLogger("movie.log")
	if err != nil {
		log.Fatalf("Failed to initialize %v", err)
	}

	defer logint.Close()

	return *logint
}

func main() {
	// Load env variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//Connect to database
	dbConnStr := os.Getenv("DB_CONNECTION_STRING")
	if dbConnStr == "" {
		log.Fatal("Database URL not set")
	}
	db, err := sql.Open("postgres", dbConnStr)
	if err != nil {
		log.Fatal("Failed to connect to database")
	}
	defer db.Close()

	// create logger
	logInstance := initializeLogger()

	//Creating handlers
	moviesHandler := handlers.MovieHanlder{
		DB: db,
	}

	http.HandleFunc("/api/movies/top", moviesHandler.GetTopMovies)
	http.HandleFunc("/api/movies/random", moviesHandler.GetRandomMovies)

	// Handle static files
	// http.Handle("/", http.FileServer(http.Dir("public")))
	addr := ":8080"
	fmt.Printf("Starting server on port: %s", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		logInstance.Error("Server failed", err)
	}
}
