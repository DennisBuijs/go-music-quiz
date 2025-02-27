package main

import (
	"log"
	"net/http"

	"dev.kipkron.music-quiz/internal/game"
	"dev.kipkron.music-quiz/internal/handlers"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", handlers.Index())
	mux.HandleFunc("POST /game/{roomID}/guess", handlers.Guess())

	songs := []game.Song{
		game.NewSong("Dancing in the Dark", "Bruce Springsteen"),
		game.NewSong("Africa", "Toto"),
		game.NewSong("Bohemian Rhapsody", "Queen"),
	}

	g := game.NewGame(songs)
	g.StartGame()

	game.Games["rock"] = g

	log.Fatal(http.ListenAndServe(":8080", mux))
}
