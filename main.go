package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	logger "github.com/cedrickewi/moviereel/Logger"
	"github.com/cedrickewi/moviereel/data"
	"github.com/cedrickewi/moviereel/handlers"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq" //import postges driver for database connection
)

func initializeLogger() *logger.Logger {
	logint, err := logger.NewLogger("movie.log")
	logint.Error("test error logging", nil)
	if err != nil {
		log.Fatalf("Failed to initialize %v", err)
	}

	defer logint.Close()

	return logint
}

func main() {
	// Load env variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// create logger
	logInstance := initializeLogger()

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

	// Init
	movieRepo, err := data.NewMovieRepository(db, logInstance)
	if err != nil {
		log.Fatalf("Fealied to initailize repository: %v", err)
	}

	// Initialize handlers
	movieHandler := handlers.NewMovieHandler(movieRepo, logInstance)

	// Set up routes
	http.HandleFunc("/api/movies/random", movieHandler.GetRandomMovies)
	http.HandleFunc("/api/movies/top/", movieHandler.GetTopMovies)
	http.HandleFunc("/api/movies/search/", movieHandler.SearchMovies)
	http.HandleFunc("/api/movies/", movieHandler.GetMovie)
	http.HandleFunc("/api/genres", movieHandler.GetGenres)
	http.HandleFunc("/api/account/register", movieHandler.GetGenres)
	http.HandleFunc("/api/account/authenticate", movieHandler.GetGenres)

	// Handle static files
	http.Handle("/", http.FileServer(http.Dir("public")))
	fmt.Println("Serving the files")

	addr := ":8080"
	fmt.Printf("Starting server on port: %s", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		logInstance.Error("Server failed", err)
	}
}
