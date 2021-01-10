package liboop

import (
	"math"
	"math/rand"
	"strconv"
)

// ComputerPlayer takes the user input
type ComputerPlayer struct {
	BasePlayer
}

func (p *ComputerPlayer) guess(playerID int) string {
	g := ""
	g = strconv.FormatInt(
		int64(rand.Intn(int(math.Pow(10, float64(p.ansSize))))-1),
		10,
	)
	p.guesses = append(p.guesses, g)
	return g
}