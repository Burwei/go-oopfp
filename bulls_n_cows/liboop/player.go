package liboop

import (
	"fmt"
)

// Player is the interface of the players
type Player interface {
	addResult([2]int8)
	guess(int) string
	getID() int8
	setID(int8)
	setAnsSize(int8)
	getGuesses() []string
	getResults() []string
}

// BasePlayer is the base class of all kind of players
type BasePlayer struct {
	id      int8
	ansSize int8
	guesses []string
	results []string
}

func (p *BasePlayer) addResult(r [2]int8) {
	result := fmt.Sprintf("%v%v%v%v", r[0], "A", r[1], "B")
	p.results = append(p.results, result)
}

func (p *BasePlayer) getID() int8 {
	return p.id
}

func (p *BasePlayer) setID(i int8){
	p.id = i
}

func (p *BasePlayer) setAnsSize(s int8){
	p.ansSize = s
}

func (p *BasePlayer) getGuesses() []string {
	return p.guesses
}

func (p *BasePlayer) getResults() []string {
	return p.results
}
