package libfp

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

/** This file contains functions that related to the game flow
 *  Use Run() to run the game service continously,
 *  and use StartNewGame to start a individual game.
 */

// GameInfo is the info of a game
type GameInfo struct {
	ansSize int8
	ans     map[byte]int8 // map[val]idx, where idx is in [0,ansSize), val is ascii
	players []func(int8) *PlayerInfo
}

// StartNewGame starts a new game
func StartNewGame() {
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
	ginfo := newGameInfo(int8(ansSize), int8(cpn))
	winner, round := gaming(&ginfo)
	if winner == 0 {
		fmt.Printf("You win the game. Takes %v rounds.\n", round)
	} else {
		fmt.Printf("Player %v win the game. Takes %v rounds.\n", winner, round)
	}
}

func newGameInfo(size int8, cpn int8) GameInfo {
	players := []func(int8) *PlayerInfo{}
	hp := NewHumanPlayerGuess(int8(0))
	players = append(players, hp)
	for i := 0; i < int(cpn); i++ {
		players = append(players, NewComputerPlayerGuess(int8(i)))
	}
	return GameInfo{
		size,
		generateAns(size),
		players,
	}
}

func generateAns(size int8) map[byte]int8 {
	rand.Seed(time.Now().UnixNano())
	ans := rand.Intn(int(math.Pow(10, float64(size)))) - 1
	ansStr := strconv.FormatInt(int64(ans), 10)
	for _, char := range ansStr {
		if strings.Count(ansStr, string(char)) > 1 {
			ans = -1
			break
		}
	}
	for ans < 0 {
		ans = rand.Intn(int(math.Pow(10, float64(size)))) - 1
		ansStr = strconv.FormatInt(int64(ans), 10)
		for _, char := range ansStr {
			if strings.Count(ansStr, string(char)) > 1 {
				ans = -1
				break
			}
		}
	}
	result := map[byte]int8{}
	for idx, char := range ansStr {
		result[byte(char)] = int8(idx)
	}
	return result
}

func getResult(ginfo *GameInfo, guess string) [2]int8 {
	result := [2]int8{}
	for i, v := range guess {
		if idx, ok := ginfo.ans[byte(v)]; ok {
			if idx == int8(i) {
				result[0]++ // bulls ++
			} else {
				result[1]++ // cows ++
			}
		}
	}
	return result
}

func gaming(ginfo *GameInfo) (int8, int16) {
	winner := int8(-1)
	round := int16(0)
	fmt.Println("------ Game starts!! ------ ")
	for {
		var yourInfo *PlayerInfo
		for id, p := range ginfo.players {
			pinfo := p(ginfo.ansSize)
			if id == 0 {
				yourInfo = pinfo
			}
			r := getResult(ginfo, pinfo.guesses[round])
			rstr := fmt.Sprintf("%v%v%v%v", r[0], "A", r[1], "B")
			pinfo.results = append(pinfo.results, rstr)
			fmt.Printf("Player %v: %vA%vB\n", id, r[0], r[1])
			if r == [2]int8{ginfo.ansSize, 0} {
				winner = int8(id)
			}
		}
		fmt.Println("Your guesses:")
		for i := int16(0); i <= round; i++ {
			fmt.Printf("%v : %v\n", yourInfo.guesses[i], yourInfo.results[i])
		}
		round++
		if winner != -1 {
			return winner, round
		}
	}
}

// Run will start running the game service
func Run() {
	reader := bufio.NewReader(os.Stdin)
	StartNewGame()
	for {
		fmt.Println("Play again?(Y/N):")
		inputStr, _ := reader.ReadString('\n')
		inputStr = inputStr[:len(inputStr)-1] // remove \n
		if inputStr == "Y" || inputStr == "y" {
			StartNewGame()
		} else if inputStr == "N" || inputStr == "n" {
			fmt.Println("Thanks for playing, bye!!")
			break
		} else {
			fmt.Println("Incorrect input")
		}
	}
}
