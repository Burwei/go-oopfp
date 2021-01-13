package liboop

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// HumanPlayer takes the user input
type HumanPlayer struct {
	BasePlayer
}

func (p *HumanPlayer) guess(size int8) string {
	g := ""
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
	g = inputStr
	p.guesses = append(p.guesses, g)
	return g
}