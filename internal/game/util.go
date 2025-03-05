package game

import (
	"fmt"
	"strings"
)

func (sched GameSchedule) String() string {
	var sb strings.Builder

	for _, state := range sched.States {
		if state.State == stateSongPlaying.State {
			sb.WriteString(fmt.Sprintf("%s\t%s\t%s\n", state.State, state.Duration, state.Song.Title))
		} else {
			sb.WriteString(fmt.Sprintf("%s\t%s\n", state.State, state.Duration))
		}
	}

	sb.WriteString("\n")

	return sb.String()
}

func (g *Game) Log() {
	fmt.Println(g.Name)
	fmt.Print(g.Schedule)
}
