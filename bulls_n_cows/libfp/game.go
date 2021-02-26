/** This file contains functions that related to the game flow
 *  Use Run() to run the game service continously,
 *  and use StartGaming to start a individual game.
 */

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

// GameInfo is the info of a game
type GameInfo struct {
	ansSize      int8                 // size of the answer
	ans          map[byte]int8        // map[val]idx, where idx is in [0,ansSize), val is ascii
	guessMethods []func(int8) *Player // players' guess methods
}

func getGameInfoFromStdin() GameInfo {
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
	return newGameInfo(int8(ansSize), int8(cpn))
}

func newGameInfo(ansSize int8, cpn int8) GameInfo {
	guesses := []func(int8) *Player{}
	guesses = append(guesses, NewHumanPlayerGuessMethod(int8(0)))
	for i := 0; i < int(cpn); i++ {
		guesses = append(guesses, NewComputerPlayerGuessMethod(int8(i)))
	}
	return GameInfo{
		ansSize,
		generateAns(ansSize),
		guesses,
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

// StartGaming starts the game loop until there's a winner
func StartGaming(ginfo GameInfo) (int8, int16) {
	winner := int8(-1)
	rounds := int16(0)
	fmt.Println("------ Game starts!! ------ ")
	for {
		var you *Player
		for id, guessAndUpdate := range ginfo.guessMethods {
			player := guessAndUpdate(ginfo.ansSize) // get updated player by taking a new guess
			if id == 0 {
				you = player
			}
			r := getResult(&ginfo, player.guesses[rounds])
			rstr := fmt.Sprintf("%v%v%v%v", r[0], "A", r[1], "B")
			player.results = append(player.results, rstr)
			fmt.Printf("Player %v: %vA%vB\n", id, r[0], r[1])
			if r == [2]int8{ginfo.ansSize, 0} {
				winner = int8(id)
			}
		}
		fmt.Println("Your guesses:")
		for i := int16(0); i <= rounds; i++ {
			fmt.Printf("%v : %v\n", you.guesses[i], you.results[i])
		}
		rounds++
		if winner != -1 {
			return winner, rounds
		}
	}
}

func newResultMsg(winner int8, rounds int16) string {
	if winner == 0 {
		return fmt.Sprintf("You win the game. Takes %v rounds.\n", rounds)
	}
	return fmt.Sprintf("Player %v win the game. Takes %v rounds.\n", winner, rounds)

}

// Run will start running the game service
func Run() {
	reader := bufio.NewReader(os.Stdin)
	ginfo := getGameInfoFromStdin()
	winner, rounds := StartGaming(ginfo)
	fmt.Printf(newResultMsg(winner, rounds))
	for {
		fmt.Println("Play again?(Y/N):")
		inputStr, _ := reader.ReadString('\n')
		inputStr = inputStr[:len(inputStr)-1] // remove \n
		if inputStr == "Y" || inputStr == "y" {
			ginfo = getGameInfoFromStdin()
			winner, rounds = StartGaming(ginfo)
			fmt.Printf(newResultMsg(winner, rounds))
		} else if inputStr == "N" || inputStr == "n" {
			fmt.Println("Thanks for playing, bye!!")
			break
		} else {
			fmt.Println("Incorrect input")
		}
	}
}
