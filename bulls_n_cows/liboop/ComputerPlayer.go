package liboop

import (
	"math"
	"math/rand"
	"strconv"
	"strings"
)

// ComputerPlayer takes the user input
type ComputerPlayer struct {
	BasePlayer
}

func (p *ComputerPlayer) guess(playerID int) string {
	ans := rand.Intn(int(math.Pow(10, float64(p.ansSize)))) - 1
	ansStr := strconv.FormatInt(int64(ans), 10)
	for _, char := range ansStr {
		if strings.Count(ansStr, string(char)) > 1 {
			ans = -1
			break
		}
	}
	for ans < 0 {
		ans = rand.Intn(int(math.Pow(10, float64(p.ansSize)))) - 1
		ansStr = strconv.FormatInt(int64(ans), 10)
		for _, char := range ansStr {
			if strings.Count(ansStr, string(char)) > 1 {
				ans = -1
				break
			}
		}
	}
	p.guesses = append(p.guesses, ansStr)
	return ansStr
}