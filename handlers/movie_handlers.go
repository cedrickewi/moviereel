package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/cedrickewi/moviereel/models"
)

type MovieHanlder struct {
	DB sql.DB
}

func (h *MovieHanlder) GetTopMovies(w http.ResponseWriter, r *http.Request) {
	movies := []models.Movie{
		{
			ID:          1,
			TMDB_ID:     181,
			Title:       "The Hacker",
			ReleaseYear: 2022,
		},
		{
			ID:          1,
			TMDB_ID:     181,
			Title:       "The Hacker",
			ReleaseYear: 2022,
		},
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(movies); err != nil {
		// TODO: Log error
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
