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
		b, _, err := reader.ReadLine()
		if err != nil {
			fmt.Println(score)
			return
		}
		s := string(b)
		part1, part2 := s[0:len(s)/2], s[len(s)/2:]
		common := findCommon(part1, part2)
		//65 97
		priority := 0
		if int(common) >= 65 && int(common) <= 90 {
			priority = int(common) - 65 + 27
		} else {
			priority = int(common) - 97 + 1
		}
		score += priority
	}
}
func findCommon(s1, s2 string) rune {
	m := map[rune]bool{}
	for _, c := range s1 {
		m[c] = true
	}
	for _, c := range s2 {
		if m[c] {
			return c
		}
	}
	return ' '
}
