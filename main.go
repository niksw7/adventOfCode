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
	tracker := map[string]bool{}
	count := 0

	reader := bufio.NewReader(inFile)
	headPosition := []int{0, 0}
	tailPosition := []int{0, 0}
	for {

		line, _, err := reader.ReadLine()
		if err != nil {
			fmt.Printf("count=%d\n", count)
			return
		}

		split := strings.Split(string(line), " ")
		character := split[0]
		move, _ := strconv.Atoi(split[1])

		for i := 1; i <= move; i++ {

			switch character {
			case "L":
				//move the head first
				headPosition = []int{headPosition[0], headPosition[1] - 1}
			case "R":
				headPosition = []int{headPosition[0], headPosition[1] + 1}
			case "U":
				headPosition = []int{headPosition[0] + 1, headPosition[1]}
			case "D":
				headPosition = []int{headPosition[0] - 1, headPosition[1]}
			}
			r, c := getPosToAdjust(headPosition, tailPosition)
			tailPosition = []int{tailPosition[0] + r, tailPosition[1] + c}
			if tracker[fmt.Sprintf("%d,%d", tailPosition[0], tailPosition[1])] {

			} else {
				tracker[fmt.Sprintf("%d,%d", tailPosition[0], tailPosition[1])] = true
				count++
			}

		}
	}

}

func getPosToAdjust(headPosition, tailPosition []int) (int, int) {
	rowDifference := headPosition[0] - tailPosition[0]
	colDifference := headPosition[1] - tailPosition[1]
	r := 0
	c := 0
	if isAdjacent(rowDifference, colDifference) {
		return 0, 0
	}

	if rowDifference > 0 {
		r = 1
	} else if rowDifference < 0 {
		r = -1
	}
	if colDifference > 0 {
		c = 1
	} else if colDifference < 0 {
		c = -1
	}
	return r, c

}

func isAdjacent(rowDifference, colDifference int) bool {
	if rowDifference == 0 && colDifference == 0 {
		return true
	} else if rowDifference == 0 && (colDifference == -1 || colDifference == 1) {
		return true
	} else if colDifference == 0 && (rowDifference == -1 || rowDifference == 1) {
		return true
	} else if colDifference+rowDifference == 0 || colDifference-rowDifference == 0 {
		return true
	}
	return false
}
