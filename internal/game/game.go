package game

import (
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
	Schedule    GameSchedule
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
	g.GenerateGameSchedule(0)
	g.Log()

	for _, state := range g.Schedule.States {
		state := state

		go func() {
			time.Sleep(time.Until(state.EndAt))
			g.Emit("gameState", state.State)
		}()
	}
}

func (g *Game) RandomSong() Song {
	return g.Songs[rand.Intn(len(g.Songs))]
}

func (g *Game) Guess(guess string) {
	if g.CurrentState().State != stateSongPlaying.State {
		return
	}

	guess = strings.ToLower(guess)

	if guess == strings.ToLower(g.CurrentSong.Title) || guess == strings.ToLower(g.CurrentSong.Artist) {
		g.Score++
	}
}

func (g *Game) Emit(eventType string, message string) {
	g.SSEServer.Publish(g.Slug, &sse.Event{
		Event: []byte(eventType),
		Data:  []byte(message),
	})
}
