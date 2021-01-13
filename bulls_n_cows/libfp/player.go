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

// PlayerInfo is the info of a player
type PlayerInfo struct {
	id      int8
	guesses []string
	results []string
}

// NewPlayerGuess creates one new guess from a specific type of player
func NewPlayerGuess(playerID int8, playerType string) func(int8) *PlayerInfo {
	pinfo := &PlayerInfo{}
	pinfo.id = playerID
	switch playerType {
	case "human":
		return func(size int8) *PlayerInfo {
			reader := bufio.NewReader(os.Stdin)
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
	default:
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
}
