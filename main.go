package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	inFile, err := os.Open("input.txt")
	if err != nil {
		panic("err" + err.Error())
	}
	result := 0
	reader := bufio.NewReader(inFile)
	for {

		for i := 0; i < 3; i++ {
			b, _, err := reader.ReadLine()
			if err != nil {
				fmt.Println(result)
				return
			}
			parts := strings.Split(string(b), ",")
			p1 := strings.Split(parts[0], "-")
			p2 := strings.Split(parts[1], "-")
			part1 := convert(p1)
			part2 := convert(p2)
			//2-8 3-7
			if Is(part1, part2) || Is(part2, part1) {
				result++
			}

		}

	}
}
func convert(p []string) []int {
	ints := []int{}
	for _, i := range p {
		elem, _ := strconv.Atoi(i)
		ints = append(ints, elem)
	}
	return ints

}

func Is(part2, part1 []int) bool {
	if part2[0] <= part1[1] && part2[0] >= part1[0] {
		return true
	}
	return false
}
