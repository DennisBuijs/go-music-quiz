package game

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
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
}

func NewGame(name string, slug string, songs []Song) *Game {
	return &Game{
		Name:  name,
		Slug:  slug,
		Songs: songs,
	}
}

func (g *Game) StartGame() {
	ticker := time.NewTicker(30 * time.Second)

	randomSong := g.Songs[rand.Intn(len(g.Songs))]
	g.CurrentSong = randomSong
	fmt.Printf("Playing %s by %s\n", randomSong.Title, randomSong.Artist)

	go func() {
		for range ticker.C {
			randomSong := g.Songs[rand.Intn(len(g.Songs))]
			g.CurrentSong = randomSong
			fmt.Printf("Playing %s by %s\n", randomSong.Title, randomSong.Artist)
		}
	}()
}

func (g *Game) Guess(guess string) {
	guess = strings.ToLower(guess)

	if guess == strings.ToLower(g.CurrentSong.Title) || guess == strings.ToLower(g.CurrentSong.Artist) {
		g.Score++
	}
}
