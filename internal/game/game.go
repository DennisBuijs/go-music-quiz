package game

import (
	"fmt"
	"math/rand"
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
	Songs       []Song
	CurrentSong Song
	Score       int
}

func NewGame(songs []Song) *Game {
	return &Game{
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
