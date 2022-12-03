package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	score := 0
	inFile, err := os.Open("input.txt")
	if err != nil {
		panic("err" + err.Error())
	}

	reader := bufio.NewReader(inFile)
	for {
		parts := make([]string, 3)
		for i := 0; i < 3; i++ {
			b, _, err := reader.ReadLine()
			if err != nil {
				fmt.Println(score)
				return
			}
			parts[i] = string(b)

		}
		common := findCommon(parts)
		priority := 0
		if int(common) >= 65 && int(common) <= 90 {
			priority = int(common) - 65 + 27
		} else {
			priority = int(common) - 97 + 1
		}
		score += priority
	}
}
func findCommon(parts []string) rune {
	m1 := map[rune]bool{}
	m2 := map[rune]bool{}
	for _, c := range parts[0] {
		m1[c] = true
	}
	for _, c := range parts[1] {
		m2[c] = true
	}

	for _, c := range parts[2] {
		if m1[c] && m2[c] {
			return c
		}
	}
	return ' '
}
