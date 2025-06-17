package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/cedrickewi/moviereel/handlers"
	"github.com/cedrickewi/moviereel/logger"
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
	// create logger
	logInstance := initializeLogger()

	moviesHandler := handlers.MovieHanlder{}

	http.HandleFunc("/api/movies/top", moviesHandler.GetTopMovies)
	http.HandleFunc("/api/movies/random", moviesHandler.GetRandomMovies)

	// Handle static files
	http.Handle("/", http.FileServer(http.Dir("public")))
	addr := ":8080"
	fmt.Printf("Starting server on port: %s", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		logInstance.Error("Server failed", err)
	}
}
