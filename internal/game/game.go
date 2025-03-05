package game

import (
	"math/rand"
	"strings"
	"time"

	"github.com/r3labs/sse/v2"
)

var Games = make(map[string]*Game)

type Song struct {
	Title    string
	Artist   string
	AudioURL string
}

func NewSong(title, artist, audioURL string) Song {
	return Song{
		Title:    title,
		Artist:   artist,
		AudioURL: audioURL,
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

func (g *Game) StartGame(amountOfSongs int) {
	g.GenerateGameSchedule(amountOfSongs)
	g.Log()

	for _, state := range g.Schedule.States {
		state := state

		go func() {
			time.Sleep(time.Until(state.EndAt))
			g.Emit("gameState", state.State)
			if state.State == stateSongPlaying.State {
				g.Emit("audio", "<audio autoplay src=\""+state.Song.AudioURL+"\">")
			}
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

	if guess == strings.ToLower(g.CurrentState().Song.Title) || guess == strings.ToLower(g.CurrentState().Song.Artist) {
		g.Score++
	}
}

func (g *Game) Emit(eventType string, message string) {
	g.SSEServer.Publish(g.Slug, &sse.Event{
		Event: []byte(eventType),
		Data:  []byte(message),
	})
}
