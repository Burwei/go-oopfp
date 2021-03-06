package liboop

import (
	"bufio"
	"fmt"
	"math"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

// Dealer type manage the procedure of the game
type Dealer struct {
	ansSize int8
	ans     map[byte]int8 // map[val]idx, where idx is in [0,ansSize), val is ascii
	players []Player
}

func (d *Dealer) init() {
	d.ansSize = 0
	d.ans = map[byte]int8{}
	d.players = []Player{}
}

// StartNewGame starts a nwe game
func (d *Dealer) StartNewGame() {
	// init the dealer
	d.init()
	// get size of the answer
	fmt.Println("Please enter the size of the answer:")
	reader := bufio.NewReader(os.Stdin)
	inputStr, _ := reader.ReadString('\n')
	inputStr = inputStr[:len(inputStr)-1] // remove \n
	ansSize, err := strconv.ParseInt(inputStr, 10, 8)
	for err != nil || ansSize <= 0 {
		fmt.Println("Incorrect number, please try again:")
		inputStr, _ = reader.ReadString('\n')
		inputStr = inputStr[:len(inputStr)-1] // remove \n
		ansSize, err = strconv.ParseInt(inputStr, 10, 8)
	}
	d.ansSize = int8(ansSize)
	playerID := int8(0)
	// add one human player
	p := &HumanPlayer{BasePlayer{}}
	p.setID(playerID)
	d.players = append(d.players, p)
	playerID++
	// add computer players
	fmt.Println("Please enter the number of computer players:")
	inputStr, _ = reader.ReadString('\n')
	inputStr = inputStr[:len(inputStr)-1] // remove \n
	cpn, err := strconv.ParseInt(inputStr, 10, 8)
	for err != nil || cpn < 0 {
		fmt.Println("Incorrect number, please try again:")
		inputStr, _ = reader.ReadString('\n')
		inputStr = inputStr[:len(inputStr)-1] // remove \n
		cpn, err = strconv.ParseInt(inputStr, 10, 8)
	}
	for i := 0; i < int(cpn); i++ {
		p := &ComputerPlayer{BasePlayer{}}
		p.setID(playerID)
		d.players = append(d.players, p)
		playerID++
	}
	d.generateAns()
	d.gaming()
}

func (d *Dealer) generateAns() {
	rand.Seed(time.Now().UnixNano())
	ans := rand.Intn(int(math.Pow(10, float64(d.ansSize)))) - 1
	ansStr := strconv.FormatInt(int64(ans), 10)
	for _, char := range ansStr {
		if strings.Count(ansStr, string(char)) > 1 {
			ans = -1
			break
		}
	}
	for ans < 0 {
		ans = rand.Intn(int(math.Pow(10, float64(d.ansSize)))) - 1
		ansStr = strconv.FormatInt(int64(ans), 10)
		for _, char := range ansStr {
			if strings.Count(ansStr, string(char)) > 1 {
				ans = -1
				break
			}
		}
	}
	for idx, char := range ansStr {
		d.ans[byte(char)] = int8(idx)
	}
}

func (d *Dealer) getResult(guess string) [2]int8 {
	result := [2]int8{}
	for i, v := range guess {
		if idx, ok := d.ans[byte(v)]; ok {
			if idx == int8(i) {
				result[0]++ // bulls ++
			} else {
				result[1]++ // cows ++
			}
		}
	}
	return result
}

func (d *Dealer) gaming() {
	win := int8(-1)
	hp := d.players[0]
	round := int16(0)
	fmt.Println("------ Game starts!! ------ ")
	for {
		for id, p := range d.players {
			r := d.getResult(p.guess(d.ansSize))
			p.addResult(r)
			fmt.Printf("Player %v: %vA%vB\n", id, r[0], r[1])
			if r == [2]int8{d.ansSize, 0} {
				win = int8(id)
			}
		}
		fmt.Println("Your guesses:")
		for i := int16(0); i <= round; i++ {
			fmt.Printf("%v : %v\n", hp.getGuesses()[i], hp.getResults()[i])
		}
		round++
		if win != -1 {
			break
		}
	}
	if win == 0 {
		fmt.Printf("You win the game. Takes %v rounds.\n", round)
	} else {
		fmt.Printf("Player %v win the game. Takes %v rounds.\n", win, round)
	}
}

// Run will start the gaming event-loop
func (d *Dealer) Run() {
	reader := bufio.NewReader(os.Stdin)
	d.StartNewGame()
	for {
		fmt.Println("Play again?(Y/N):")
		inputStr, _ := reader.ReadString('\n')
		inputStr = inputStr[:len(inputStr)-1] // remove \n
		if inputStr == "Y" || inputStr == "y" {
			d.StartNewGame()
		} else if inputStr == "N" || inputStr == "n" {
			fmt.Println("Thanks for playing, bye!!")
			break
		} else {
			fmt.Println("Incorrect input")
		}
	}
}
