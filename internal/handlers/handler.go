package handlers

import (
	"html/template"
	"net/http"

	"dev.kipkron.music-quiz/internal/game"
)

func Index() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var rooms []RoomResponse

		for _, room := range game.Games {
			room := RoomResponse{
				Name: room.Name,
				Slug: room.Slug,
			}
			rooms = append(rooms, room)
		}

		response := RoomsIndexResponse{
			Title: "Pick a room",
			Rooms: rooms,
		}

		tmpl := template.Must(template.ParseFiles("public/index.html", "public/rooms_index.html"))
		err := tmpl.Execute(w, response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func Room() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		g, ok := game.Games[r.PathValue("roomID")]
		if !ok {
			http.Error(w, "Game not found", http.StatusNotFound)
			return
		}

		response := RoomsShowResponse{
			Name:  g.Name,
			Slug:  g.Slug,
			State: g.CurrentState(),
			Score: g.Score,
		}

		tmpl := template.Must(template.ParseFiles("public/index.html", "public/rooms_show.html"))
		err := tmpl.Execute(w, response)
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

		g.Guess(guess)

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

type RoomResponse struct {
	Name string
	Slug string
}

type RoomsIndexResponse struct {
	Title string
	Rooms []RoomResponse
}

type RoomsShowResponse struct {
	Name  string
	Slug  string
	State game.GameState
	Score int
}

type GuessResponse struct {
	Score int
}
