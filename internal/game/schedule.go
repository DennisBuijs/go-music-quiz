package game

import "time"

type GameState struct {
	State    string
	Duration time.Duration
	Song     Song
	EndAt    time.Time
}

type GameSchedule struct {
	States []GameState
}

var stateGameStarting = GameState{
	State:    "starting",
	Duration: 15 * time.Second,
}

var stateSongPlaying = GameState{
	State:    "songPlaying",
	Duration: 30 * time.Second,
}

var stateSongBreak = GameState{
	State:    "songBreak",
	Duration: 15 * time.Second,
}

var stateGameEnding = GameState{
	State:    "gameEnding",
	Duration: 30 * time.Second,
}

var stateGameEnded = GameState{
	State:    "gameEnded",
	Duration: 0,
}

func (g *Game) GenerateGameSchedule(amountOfSongs int) {
	schedule := GameSchedule{
		States: make([]GameState, 2+(amountOfSongs*2)),
	}

	schedule.States[0] = GameState{
		State:    stateGameStarting.State,
		Duration: stateGameStarting.Duration,
		EndAt:    time.Now().Add(stateGameStarting.Duration),
	}

	for i := 1; i < (amountOfSongs * 2); i += 2 {
		schedule.States[i] = GameState{
			State:    stateSongPlaying.State,
			Duration: stateSongPlaying.Duration,
			EndAt:    schedule.States[i-1].EndAt.Add(stateSongPlaying.Duration),
			Song:     g.RandomSong(),
		}
		schedule.States[i+1] = GameState{
			State:    stateSongBreak.State,
			Duration: stateSongBreak.Duration,
			EndAt:    schedule.States[i].EndAt.Add(stateSongBreak.Duration),
		}
	}

	schedule.States[1+(amountOfSongs*2)] = GameState{
		State:    stateGameEnding.State,
		Duration: stateGameEnding.Duration,
		EndAt:    schedule.States[amountOfSongs*2].EndAt.Add(stateGameEnding.Duration),
	}

	g.Schedule = schedule
}

func (g *Game) CurrentState() GameState {
	for _, state := range g.Schedule.States {
		if state.EndAt.After(time.Now()) {
			return state
		}
	}

	return stateGameEnded
}
