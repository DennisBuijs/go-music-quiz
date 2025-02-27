package game

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/r3labs/sse/v2"
)

var Games = make(map[string]*Game)

type Song struct {
	Title  string
	Artist string
}

func NewSong(title, artist string) Song {
	return Song{
		Title:  title,
		Artist: artist,
	}
}

type Game struct {
	Name        string
	Slug        string
	Songs       []Song
	CurrentSong Song
	Score       int
	SSEServer   *sse.Server
}

func NewGame(name string, slug string, songs []Song, sses *sse.Server) *Game {
	return &Game{
		Name:      name,
		Slug:      slug,
		Songs:     songs,
		SSEServer: sses,
	}
}

func (g *Game) StartGame() {
	ticker := time.NewTicker(30 * time.Second)

	randomSong := g.Songs[rand.Intn(len(g.Songs))]
	g.CurrentSong = randomSong
	message := fmt.Sprintf("<div>Playing %s by %s</div>", randomSong.Title, randomSong.Artist)
	g.Emit("nextSong", message)

	go func() {
		for range ticker.C {
			randomSong := g.Songs[rand.Intn(len(g.Songs))]
			g.CurrentSong = randomSong

			message := fmt.Sprintf("<div>Playing %s by %s</div>", randomSong.Title, randomSong.Artist)
			g.Emit("nextSong", message)
		}
	}()
}

func (g *Game) Emit(eventType string, message string) {
	g.SSEServer.Publish(g.Slug, &sse.Event{
		Event: []byte(eventType),
		Data:  []byte(message),
	})
}

func (g *Game) Guess(guess string) {
	guess = strings.ToLower(guess)

	if guess == strings.ToLower(g.CurrentSong.Title) || guess == strings.ToLower(g.CurrentSong.Artist) {
		g.Score++
	}
}
