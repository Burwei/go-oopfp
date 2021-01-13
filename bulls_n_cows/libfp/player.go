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

/** This file contains functions that related to the player.
 *  Use NewXXXXXGuess() to guess a number according to the player type.
 *  Player's info is stored in the closure.
 */

// PlayerInfo is the info of a player
type PlayerInfo struct {
	id      int8
	guesses []string
	results []string
}

// NewHumanPlayerGuess creates a human player.
func NewHumanPlayerGuess(playerID int8) func(int8) *PlayerInfo {
	pinfo := &PlayerInfo{}
	pinfo.id = playerID
	reader := bufio.NewReader(os.Stdin)
	return func(size int8) *PlayerInfo {
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
		for err != nil || int8(len(inputStr)) != size || n < 0 {
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
		pinfo.guesses = append(pinfo.guesses, inputStr)
		return pinfo
	}
}

// NewComputerPlayerGuess creates a computer player.
func NewComputerPlayerGuess(playerID int8) func(int8) *PlayerInfo {
	pinfo := &PlayerInfo{}
	pinfo.id = playerID
	return func(size int8) *PlayerInfo {
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
		pinfo.guesses = append(pinfo.guesses, ansStr)
		return pinfo
	}
}
