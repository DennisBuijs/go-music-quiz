package handlers

import (
	"html/template"
	"net/http"
	"strings"

	"dev.kipkron.music-quiz/internal/game"
)

func Index() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("public/index.html"))
		err := tmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func Guess() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		g, ok := game.Games[r.PathValue("roomID")]
		if !ok {
			http.Error(w, "Game not found", http.StatusNotFound)
			return
		}

		guess := r.FormValue("guess")
		if guess == "" {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		guess = strings.ToLower(guess)

		if guess == strings.ToLower(g.CurrentSong.Title) || guess == strings.ToLower(g.CurrentSong.Artist) {
			g.Score++
		}

		response := GuessResponse{
			Score: g.Score,
		}

		tmpl := template.Must(template.ParseFiles("public/scoreboard.html"))
		err := tmpl.Execute(w, response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

type GuessResponse struct {
	Score int
}
