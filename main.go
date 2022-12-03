package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	score := 0
	inFile, err := os.Open("input.txt")
	if err != nil {
		panic("err" + err.Error())
	}

	reader := bufio.NewReader(inFile)
	for {
		b, _, err := reader.ReadLine()
		if err != nil {
			fmt.Println(score)
			return
		}
		words := strings.Split(string(b), " ")

		switch words[1] {
		case "X":
			//lloose
			words[1] = outcome(words[0], false)

		case "Y": //draw
			words[1] = mapIt(words[0])
		case "Z": //win
			words[1] = outcome(words[0], true)
		}

		score += play(words[0], words[1])
	}
}
func mapIt(word string) string {
	switch word {
	case "A":
		return "X"
	case "B":
		return "Y"
	case "C":
		return "Z"
	}
	return ""
}

// A=Rock
// B=Paper
// C=Scissor
func outcome(play string, shouldWin bool) string {
	switch play {
	case "A":
		if shouldWin {
			return "Y"
		} else {
			return "Z"
		}
	case "B":
		if shouldWin {
			return "Z"
		} else {
			return "X"
		}
	case "C":
		if shouldWin {
			return "X"
		} else {
			return "Y"
		}
	}
	panic("ss")

}
func play(player1 string, player2 string) int {
	switch player1 + player2 {
	case "AX":
		return 1 + 3
	case "AY":
		return 2 + 6
	case "AZ":
		return 3
	case "BX":
		return 1
	case "BY":
		return 2 + 3
	case "BZ":
		return 3 + 6
	case "CX":
		return 1 + 6
	case "CY":
		return 2
	case "CZ":
		return 3 + 3

	}
	return 0
}
