package main

import (
	"log"
	"net/http"
	"strings"

	"dev.kipkron.music-quiz/internal/game"
	"dev.kipkron.music-quiz/internal/handlers"
	"github.com/r3labs/sse/v2"
)

func main() {
	sses := sse.New()
	sses.BufferSize = 0
	sses.AutoReplay = false

	startClassicRockRoom(sses)
	startClassicPopRoom(sses)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /", handlers.Index())
	mux.HandleFunc("GET /room/{roomID}", handlers.Room())
	mux.HandleFunc("POST /room/{roomID}/guess", handlers.Guess())

	mux.HandleFunc("GET /assets/", assetHandler())
	mux.HandleFunc("GET /sse", sses.ServeHTTP)

	log.Fatal(http.ListenAndServe("0.0.0.0:8080", mux))
}

func startClassicRockRoom(sses *sse.Server) {
	songs := []game.Song{
		game.NewSong("Dancing in the Dark", "Bruce Springsteen", "someURL"),
		game.NewSong("Africa", "Toto", "someURL"),
		game.NewSong("Bohemian Rhapsody", "Queen", "someURL"),
	}

	g := game.NewGame("Classic Rock", "classic-rock", songs, sses)
	g.StartGame(3)

	game.Games[g.Slug] = g

	sses.CreateStream(g.Slug)
}

func startClassicPopRoom(sses *sse.Server) {
	songs := []game.Song{
		game.NewSong("I Will Always Love You", "Whitney Houston", "someURL"),
		game.NewSong("Billie Jean", "Michael Jackson", "someURL"),
		game.NewSong("Shape of You", "Ed Sheeran", "someURL"),
	}

	g := game.NewGame("Classic Pop", "classic-pop", songs, sses)
	g.StartGame(3)

	game.Games[g.Slug] = g

	sses.CreateStream(g.Slug)
}

func assetHandler() http.HandlerFunc {
	fs := http.FileServer(http.Dir("./public"))
	return func(w http.ResponseWriter, r *http.Request) {
		r.URL.Path = strings.TrimPrefix(r.URL.Path, "/assets")
		fs.ServeHTTP(w, r)
	}
}
