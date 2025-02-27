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
	mux.HandleFunc("GET /room/{roomID}", handlers.Room())
	mux.HandleFunc("POST /room/{roomID}/guess", handlers.Guess())

	startClassicRockRoom()
	startClassicPopRoom()

	log.Fatal(http.ListenAndServe(":8080", mux))
}

func startClassicRockRoom() {
	songs := []game.Song{
		game.NewSong("Dancing in the Dark", "Bruce Springsteen"),
		game.NewSong("Africa", "Toto"),
		game.NewSong("Bohemian Rhapsody", "Queen"),
	}

	g := game.NewGame("Classic Rock", "classic-rock", songs)
	g.StartGame()

	game.Games[g.Slug] = g
}

func startClassicPopRoom() {
	songs := []game.Song{
		game.NewSong("I Will Always Love You", "Whitney Houston"),
		game.NewSong("Billie Jean", "Michael Jackson"),
		game.NewSong("Shape of You", "Ed Sheeran"),
	}

	g := game.NewGame("Classic Pop", "classic-pop", songs)
	g.StartGame()

	game.Games[g.Slug] = g
}
