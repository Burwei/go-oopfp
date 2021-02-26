/** This file contains functions that related to the player.
 *  Use NewXXXXXGuessMethod() to guess a number according to the player type.
 *  Player is stored in the closure.
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
)

// Player contains the info of a player
type Player struct {
	id      int8
	guesses []string
	results []string
}

// NewHumanPlayerGuessMethod creates a human player's guess method.
// The guess will returns a updated player.
func NewHumanPlayerGuessMethod(id int8) func(int8) *Player {
	player := &Player{}
	player.id = id
	reader := bufio.NewReader(os.Stdin)
	return func(anslen int8) *Player {
		fmt.Println("Please enter a number:")
		inputStr, _ := reader.ReadString('\n')
		inputStr = inputStr[:len(inputStr)-1] // remove \n
		n, err := strconv.ParseInt(inputStr, 10, 64)
		for _, char := range inputStr {
			if strings.Count(inputStr, string(char)) > 1 {
				n = -1
				break
			}
		}
		for err != nil || int8(len(inputStr)) != anslen || n < 0 {
			fmt.Println("The number is incorrect, please try again:")
			inputStr, _ = reader.ReadString('\n')
			inputStr = inputStr[:len(inputStr)-1] // remove \n
			n, err = strconv.ParseInt(inputStr, 10, 64)
			for _, char := range inputStr {
				if strings.Count(inputStr, string(char)) > 1 {
					n = -1
					break
				}
			}
		}
		player.guesses = append(player.guesses, inputStr)
		return player
	}
}

// NewComputerPlayerGuessMethod creates a computer player's guess method.
// The guess will returns a updated player.
func NewComputerPlayerGuessMethod(id int8) func(int8) *Player {
	player := &Player{}
	player.id = id
	return func(anslen int8) *Player {
		ans := rand.Intn(int(math.Pow(10, float64(anslen)))) - 1
		ansStr := strconv.FormatInt(int64(ans), 10)
		for _, char := range ansStr {
			if strings.Count(ansStr, string(char)) > 1 {
				ans = -1
				break
			}
		}
		for ans < 0 {
			ans = rand.Intn(int(math.Pow(10, float64(anslen)))) - 1
			ansStr = strconv.FormatInt(int64(ans), 10)
			for _, char := range ansStr {
				if strings.Count(ansStr, string(char)) > 1 {
					ans = -1
					break
				}
			}
		}
		player.guesses = append(player.guesses, ansStr)
		return player
	}
}
